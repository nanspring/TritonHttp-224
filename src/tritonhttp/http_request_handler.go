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

	req_header := HttpRequestHeader{Host:"",Connection:""}
	res_header := HttpResponseHeader{ResponseCode: "", Server: SERVER_NAME, LastModified: "", ContentType: "", ContentLength: "", Connection:"open", FilePath: ""}
	remaining := ""
	buf := make([]byte,1024)
	var close bool

	// set read reqeust initial line to false
	read_initial := false


	for {

		//set times out before each read
		err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))

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
				return 
			}
			if err1, ok := err.(net.Error); ok && err1.Timeout() {
				log.Println("Timeout", err1)
				log.Println("connection close")
				conn.Close()
				return
            }
		}

		data  := buf[:size]
		remaining += string(data)

		//check if remaining contains full request. If not, connection wait to read again before timesout
		for strings.Contains(remaining, DELIMITER){
			i := strings.Index(remaining,DELIMITER)
			line := remaining[:i] // get the full line

			//parse initial line
			if !read_initial{
				if !hs.ExamParseInitalLine(line, &res_header, conn){
				}
				read_initial = true

			}else if len(line) != 0{ // parse request key value pair 
				if !hs.ParseKeyValuePair(line, &req_header,conn){
				}
			}else{  //line is empty, meaning it is the end of full request, return response
				close = hs.sendResponse(&req_header, &res_header, conn)
				if close{
					return
				}
			}
			remaining = remaining [i+2:]
		}
	}
}
