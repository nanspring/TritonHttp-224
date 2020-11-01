package tritonhttp

import (
	"log"
	"net"
)
/** 
	Initialize the tritonhttp server by populating HttpServer structure
**/
func NewHttpdServer(port, docRoot, mimePath string) (*HttpServer, error) {
	//panic("todo - NewHttpdServer")
	// Initialize mimeMap for server to refer
	// Return pointer to HttpServer

	log.Println("Server has doc root as:", docRoot)
	log.Println("Server has mime types file at:", mimePath)

	// read mimeMap file to a map
	mimeMap,err := ParseMIME(mimePath)

	hs := HttpServer{ServerPort: port, DocRoot: docRoot, MIMEPath: mimePath, MIMEMap: mimeMap}
	return &hs, err
}

/** 
	Start the tritonhttp server
**/
func (hs *HttpServer) Start() (err error) {
	//panic("todo - StartServer")

	// Start listening to the server port

	// Accept connection from client

	// Spawn a go routine to handle request
	port := hs.ServerPort
	host := "0.0.0.0"
	//delim := "/r/n"
	ln, err := net.Listen("tcp", host+port)
	defer ln.Close()
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Listening to connections at '"+host+"' on port", port)

	for {
		conn, err := ln.Accept()
		if err != nil{
			log.Panicln(err)
		}
		
		go hs.handleConnection(conn)
	}

	return err


}

