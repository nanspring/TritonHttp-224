package tritonhttp

import (
	"bufio"
	"os"
	"strings"
	"net"
	"log"

)
/** 
	Load and parse the mime.types file 
**/
func ParseMIME(MIMEPath string) (MIMEMap map[string]string, err error) {
	//panic("todo - ParseMIME")

	MIMEMap = make(map[string]string)
	file, err := os.Open(MIMEPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		parts := strings.Split(scanner.Text()," ")
		MIMEMap[parts[0]] = parts[1]
	}
	return MIMEMap,err
}

/**
	exam and parse request header initial line
	200 means no error
	400 means bad request
	404 means file not found
**/
func (hs *HttpServer) ExamParseInitalLine(inital_line string, req_header *HttpRequestHeader, conn net.Conn) (code int){
	code = 200
	parts := strings.Split(inital_line, " ")
	if len(parts) != 3 || parts[0] != "GET" || parts[2] != "HTTP/1.1"{
		hs.handleBadRequest(conn)
		return 400
	}
	filePath := hs.DocRoot + parts[1]
	if !fileExists(filePath){
		hs.handleFileNotFoundRequest(req_header,conn)
		return 404
	}
	
	req_header.URL = parts[1]
	return code

}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}



