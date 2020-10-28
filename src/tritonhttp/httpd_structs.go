package tritonhttp

type HttpServer	struct {
	ServerPort	string
	DocRoot		string
	MIMEPath	string
	MIMEMap		map[string]string
}

type HttpResponseHeader struct {
	// Add any fields required for the response here
	Server string
	LastModified string
	ContentType string
	ContentLength uint64
	Connection string
}

type HttpRequestHeader struct {
	// Add any fields required for the request here
	Host string
	Connection string
}