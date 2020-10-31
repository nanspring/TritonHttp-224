package tritonhttp

import (
	"net"
)


func (hs *HttpServer) handleBadRequest(conn net.Conn) {
	//panic("todo - handleBadRequest")
	response := "HTTP/1.1 400 Bad Request"
	response += "Server: "+SERVER_NAME+DELIMITER
	bResponse := []byte(response)
	conn.Write(bResponse)
}

func (hs *HttpServer) handleFileNotFoundRequest(requestHeader *HttpRequestHeader, conn net.Conn) {
	//panic("todo - handleFileNotFoundRequest")

	if len(requestHeader.Host) == 0 {
		hs.handleBadRequest(conn)
		return
	}
	response := "HTTP/1.1 404 Not Found"+DELIMITER
	response += "Server: "+SERVER_NAME+DELIMITER
	bResponse := []byte(response)
	conn.Write(bResponse)
}

func (hs *HttpServer) handleResponse(requestHeader *HttpRequestHeader, conn net.Conn) (result string) {
	panic("todo - handleResponse")
}

func (hs *HttpServer) sendResponse(responseHeader HttpResponseHeader, conn net.Conn) {
	panic("todo - sendResponse")

	// Send headers

	// Send file if required

	// Hint - Use the bufio package to write response
}
