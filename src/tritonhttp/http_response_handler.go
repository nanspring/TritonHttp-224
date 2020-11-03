package tritonhttp

import (
	"net"
	"log"
	"io"
	"os"
	"bufio"
)
const BUFFERSIZE int = 1024

func (hs *HttpServer) handleBadRequest(conn net.Conn) {
	//panic("todo - handleBadRequest")
	response := "HTTP/1.1 400 Bad Request" + DELIMITER
	response += "Server: "+SERVER_NAME + DELIMITER
	response += DELIMITER
	bResponse := []byte(response)
	conn.Write(bResponse)
	conn.Close()
}

func (hs *HttpServer) handleFileNotFoundRequest(conn net.Conn) {
	//panic("todo - handleFileNotFoundRequest")
	response := "HTTP/1.1 404 Not Found" + DELIMITER
	response += "Server: " + SERVER_NAME+DELIMITER
	response += DELIMITER
	bResponse := []byte(response)
	conn.Write(bResponse)
}

func (hs *HttpServer) handleResponse(req_header *HttpRequestHeader, responseHeader *HttpResponseHeader, conn net.Conn) (result string) {
	//panic("todo - handleResponse")
	result += "HTTP/1.1 "+responseHeader.ResponseCode+DELIMITER
	result += "Server: " + responseHeader.Server + DELIMITER
	result += "Last-Modified: " + responseHeader.LastModified + DELIMITER
	result += "Content-Type: " + responseHeader.ContentType + DELIMITER
	result += "Content-Length: " + responseHeader.ContentLength + DELIMITER
	if req_header.Connection == "close"{
		result += "Connection: close" +DELIMITER
	}
	result += DELIMITER
	return result

}

/**
	send response
	send file if required
**/
func (hs *HttpServer) sendResponse(req_header *HttpRequestHeader, responseHeader *HttpResponseHeader, conn net.Conn) bool{

	if len(req_header.Host) == 0{
		log.Println("br 1")
		hs.handleBadRequest(conn)
		return false
	}

	if responseHeader.ResponseCode == "404"{
		log.Println("fnf 1", responseHeader.FilePath)
		hs.handleFileNotFoundRequest(conn)
		return false
	}
	response := hs.handleResponse(req_header, responseHeader, conn)
	bResponse := []byte(response)
	conn.Write(bResponse)

	// read requesting file
	file, err := os.Open(responseHeader.FilePath)
	defer file.Close()
	if err != nil {
		log.Println("file path", responseHeader.FilePath)
		log.Println("open body file error", err)
		return false
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, BUFFERSIZE)
	for{
		size, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		conn.Write(buffer[:size])
	}
	if req_header.Connection == "close"{
		log.Println("clinet send close request, close connection")
		conn.Close()
		return true
	}else{

		//clear req_header and res_header for new pipeline request
		req_header.Host = ""
		req_header.Connection = ""
		responseHeader.ResponseCode = ""
		responseHeader.LastModified = "" 
		responseHeader.ContentType = ""
		responseHeader.ContentLength = ""
		responseHeader.Connection = "open"
		responseHeader.FilePath = ""
		return false
	}
}
