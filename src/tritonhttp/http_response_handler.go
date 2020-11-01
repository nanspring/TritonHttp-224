package tritonhttp

import (
	"net"
)


func (hs *HttpServer) handleBadRequest(conn net.Conn) {
	//panic("todo - handleBadRequest")
	response := "HTTP/1.1 400 Bad Request" + DELIMITER
	response += "Server: "+SERVER_NAME + DELIMITER
	bResponse := []byte(response)
	conn.Write(bResponse)
}

func (hs *HttpServer) handleFileNotFoundRequest(requestHeader *HttpRequestHeader, conn net.Conn) {
	//panic("todo - handleFileNotFoundRequest")

	if len(requestHeader.Host) == 0 {
		hs.handleBadRequest(conn)
		return
	}
	response := "HTTP/1.1 404 Not Found" + DELIMITER
	response += "Server: " + SERVER_NAME+DELIMITER
	bResponse := []byte(response)
	conn.Write(bResponse)
}

func (hs *HttpServer) handleResponse(responseHeader *HttpResponseHeader, conn net.Conn) (result string) {
	//panic("todo - handleResponse")
	result += "HTTP/1.1 "+responseHeader.ResponseCode+DELIMITER
	result += "Server: " + responseHeader.Server + DELIMITER
	result += "Last-Modified: " + responseHeader.LastModified + DELIMITER
	result += "Content-Type: " + responseHeader.ContentType + DELIMITER
	result += "Content-Length: " + responseHeader.ContentLength + DELIMITER
	result += DELIMITER
	return result

}

func (hs *HttpServer) sendResponse(res_header *HttpRequestHeader, responseHeader *HttpResponseHeader, conn net.Conn) {
	//panic("todo - sendResponse")

	// Send headers

	// Send file if required

	// Hint - Use the bufio package to write response

	if responseHeader.ResponseCode == "404"{
		hs.handleFileNotFoundRequest(res_header, conn)
		return
	}
	response := hs.handleResponse(responseHeader, conn)
	bResponse := []byte(response)
	conn.Write(bResponse)
}
