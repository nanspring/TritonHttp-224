package tritonhttp

type HttpServer	struct {
	ServerPort	string
	DocRoot		string
	MIMEPath	string
	MIMEMap		map[string]string
}

type HttpResponseHeader struct {
	// Add any fields required for the response here
	ResponseCode string
	Server string
	LastModified string
	ContentType string
	ContentLength string
	Connection string
	FilePath string
}

type HttpRequestHeader struct {
	// Add any fields required for the request here
	Host string
	Connection string
}