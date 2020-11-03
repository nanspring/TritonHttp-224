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
	bResponse := []byte(response)
	conn.Write(bResponse)
	conn.Close()
}

func (hs *HttpServer) handleFileNotFoundRequest(conn net.Conn) {
	//panic("todo - handleFileNotFoundRequest")
	response := "HTTP/1.1 404 Not Found" + DELIMITER
	response += "Server: " + SERVER_NAME+DELIMITER
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

func (hs *HttpServer) sendResponse(req_header *HttpRequestHeader, responseHeader *HttpResponseHeader, conn net.Conn) bool{
	//panic("todo - sendResponse")

	// Send headers

	// Send file if required

	// Hint - Use the bufio package to write response
	if responseHeader.ResponseCode == "404"{
		hs.handleFileNotFoundRequest(conn)
		return false
	}
	response := hs.handleResponse(req_header, responseHeader, conn)
	bResponse := []byte(response)
	conn.Write(bResponse)
	file, err := os.Open(responseHeader.FilePath)
	defer file.Close()
	if err != nil {
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
		return false
	}
}
