package tritonhttp

import (
	"os"
	"bufio"
	"log"
	"strings"
)
/** 
	Initialize the tritonhttp server by populating HttpServer structure
**/
func NewHttpdServer(port, docRoot, mimePath string) (*HttpServer, error) {
	//panic("todo - NewHttpdServer")

	// Initialize mimeMap for server to refer

	// Return pointer to HttpServer
	log.Println("here!")
	mimeMap := make(map[string]string)
	file, err := os.Open(mimePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		parts := strings.Split(scanner.Text()," ")
		mimeMap[parts[0]] = parts[1]
	}
	//log.Println(mimeMap)
	


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

