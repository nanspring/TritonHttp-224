package tritonhttp

import (
	"os"
	"encoding/csv"
	"log"
)
/** 
	Initialize the tritonhttp server by populating HttpServer structure
**/
func NewHttpdServer(port, docRoot, mimePath string) (*HttpServer, error) {
	//panic("todo - NewHttpdServer")

	// Initialize mimeMap for server to refer

	// Return pointer to HttpServer
	mimeMap := make(map[string]string)
	file, err := os.Open(mimePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	csv_reader := csv.NewReader(file)
	if parts, err := csv_reader.Read(); err == nil {
		log.Println(parts)
		mimeMap[parts[0]] = parts[1]
	} else {
		log.Fatal(err)
	}


	http := HttpServer{ServerPort: port, DocRoot: docRoot, MIMEPath: mimePath, MIMEMap: mimeMap}
	return &http,err
}

/** 
	Start the tritonhttp server
**/
func (hs *HttpServer) Start() (err error) {
	panic("todo - StartServer")

	// Start listening to the server port

	// Accept connection from client

	// Spawn a go routine to handle request

}

