package tritonhttp

import (
	"bufio"
	"os"
	"strings"
	"log"
	"net"
	"strconv"
	"time"
	"path"
	"path/filepath"
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
func (hs *HttpServer) ExamParseInitalLine(inital_line string,res_header *HttpResponseHeader, conn net.Conn){
	parts := strings.Split(inital_line, " ")
	if len(parts) != 3 || parts[0] != "GET" || parts[2] != "HTTP/1.1"{
		res_header.ResponseCode = "400"
		hs.handleBadRequest(conn)
		return
	}

	url := parts[1]
	if(url == "/"){
		url = "/index.html"
	}

	filePath := path.Join(hs.DocRoot, url)
	if !fileExists(filePath,res_header){
		res_header.ResponseCode = "404"
		return
	}
	// check whether escape the doc root
	fileAbsPath, _ := filepath.Abs(filePath)
	RootPath, _ := filepath.Abs(hs.DocRoot)
	matched := strings.Contains(fileAbsPath, RootPath)
	if !matched{
		res_header.ResponseCode = "404"
		return
	}
	mType := fileAbsPath[strings.Index(fileAbsPath,"."):]
	if val, ok := hs.MIMEMap[mType]; ok{
		res_header.ContentType = val
	}else{
		log.Println("Unkown MIME type")
		res_header.ContentType = "application/octet-stream"
	}
	res_header.ResponseCode = "200 OK"
}

/**
	check request key value pair
	put key value pair to request header
**/
func (hs *HttpServer) ParseKeyValuePair(input string, req_header *HttpRequestHeader, conn net.Conn){
	if strings.Contains(input,":"){
		parts := strings.Split(input, ":")

		if len(parts) != 2 {
			hs.handleBadRequest(conn)
		}else{

			key := parts[0]
			// trim leading and tailing space of "value"
			value := strings.TrimSpace(parts[1])

			if key == "Host"{
				req_header.Host = value
			}else if key == "Connection"{
				req_header.Connection = value
			}
		}
	}else{
		hs.handleBadRequest(conn)
	}
	if len(req_header.Host) == 0 {
		hs.handleBadRequest(conn)
		return
	}
}

func fileExists(filename string,res_header *HttpResponseHeader) bool {
    file, err := os.Stat(filename)
    if os.IsNotExist(err) {
		return false
	}else if !file.IsDir(){
		filesize := strconv.FormatInt(file.Size(),10)
		res_header.LastModified = file.ModTime().Format(time.RFC850)
		res_header.ContentLength = filesize
		res_header.FilePath = filename
		return true
	}else{
		return false
	}
}
