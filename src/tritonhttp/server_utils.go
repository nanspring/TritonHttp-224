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
func (hs *HttpServer) ExamParseInitalLine(inital_line string,res_header *HttpResponseHeader, conn net.Conn) bool {
	parts := strings.Split(inital_line, " ")
	if len(parts) != 3 || parts[0] != "GET" || parts[2] != "HTTP/1.1"{
		res_header.ResponseCode = "400"
		hs.handleBadRequest(conn)
		return false
	}
	url := parts[1]
	if (len(url)==0 || url[0] != '/'){
		res_header.ResponseCode = "400"
		hs.handleBadRequest(conn)
		return false
	}
	if(url == "/"){
		url = "/index.html"
	}
	fileAbsPath, _ := filepath.Abs(path.Join(hs.DocRoot, url))
	rootAbsPath, _ := filepath.Abs(hs.DocRoot)
	if !fileValid(fileAbsPath, rootAbsPath, res_header){
		res_header.ResponseCode = "404"
		return false
	}
	fileAbsPath = res_header.FilePath
	mType := fileAbsPath[strings.Index(fileAbsPath,"."):]
	if val, ok := hs.MIMEMap[mType]; ok{
		res_header.ContentType = val
	}else{
		log.Println("Unkown MIME type")
		res_header.ContentType = "application/octet-stream"
	}
	res_header.ResponseCode = "200 OK"
	return true
}

/**
	check request key value pair
	put key value pair to request header
**/
func (hs *HttpServer) ParseKeyValuePair(input string, req_header *HttpRequestHeader, conn net.Conn) bool{
	if strings.Contains(input,":"){
		parts := strings.Split(input, ":")

		if len(parts) != 2 {
			hs.handleBadRequest(conn)
			return false
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
		return false
	}
	// if len(req_header.Host) == 0{
	// 	hs.handleBadRequest(conn)
	// 	return false
	// }
	return true
}

func fileValid(filename string, root string, res_header *HttpResponseHeader) bool {
	file, err := os.Stat(filename)
	if os.IsNotExist(err){
		return false
	}else if !strings.Contains(filename, root){
		return false
	}else if file.IsDir(){
		filename = path.Join(filename, "index.html")
		file, _ = os.Stat(filename)
		filesize := strconv.FormatInt(file.Size(),10)
		res_header.LastModified = file.ModTime().Format(time.RFC850)
		res_header.ContentLength = filesize
		res_header.FilePath = filename
		return true	
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
