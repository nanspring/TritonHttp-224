package tritonhttp

import (
	"net"
	"strings"
	"time"
	"log"
)

// Delimiter and Server name
const  DELIMITER string = "\r\n"
const SERVER_NAME string = "Go-Triton-Server-1.0"

/* 
For a connection, keep handling requests until 
	1. a timeout occurs or
	2. client closes connection or
	3. client sends a bad request
*/
func (hs *HttpServer) handleConnection(conn net.Conn) {

	//panic("todo - handleConnection")

	// Start a loop for reading requests continuously
	// Set a timeout for read operation
	// Read from the connection socket into a buffer
	// Validate the request lines that were read
	// Handle any complete requests
	// Update any ongoing requests
	// If reusing read buffer, truncate it before next read

	req_header := HttpRequestHeader{Host:"", URL:"", Connection:"open"}
	remaining := ""
	buf := make([]byte,1024)
	for {

		//set times out before each read
		err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		
		// initial line
		read_initial := false

		if err != nil {
			log.Println("set times out fail :",err)
			return
		}

		// connection read request into buffer
		size, err := conn.Read(buf)

		//if times out appear, need to close connection
		if err != nil{
			//if there is incomplete request, it is a bad request
			if len(remaining) > 0{
				hs.handleBadRequest(conn)
			}
			conn.Close()
			return
		}

		data  := buf[:size]
		remaining += string(data)

		//check if remaining contains full request. If not, connection wait to read again before timesout
		for strings.Contains(remaining,DELIMITER){
			// todo
			i := strings.Index(remaining,DELIMITER)
			log.Println("remain: ",remaining)
			//parse initial line
			if !read_initial{
				code := hs.ExamParseInitalLine(remaining[:i],&req_header, conn)
				log.Println(code)
				log.Println("after: ",remaining)
				read_initial = true
			}
			remaining = remaining [i+2:]
		}
	}
}
