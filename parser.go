package main

import (
	"strings"
	"path"
)

type HTTPReq struct {
	method, resource, version string
	host                      string
	userAgent string
	accept []string
	authorization string
}

// Parse string HTTP request and construct a HTTPReq struct with
// HTTP request info
func ParseHTTPRequest(request string) HTTPReq {
	lines := strings.Split(request, "\n")
	var req HTTPReq

	// method, resource, version
	head := strings.Split(lines[0], " ")
	req.method = head[0]
	req.resource = path.Clean(head[1])
	req.version = head[2]

	// http headers
	headers := lines[1:len(lines)-1] // last element is blank
	for _, line := range headers {
		l := strings.SplitN(line, ": ", 2)
		header := l[0]
		value := l[1]
		switch header {
			case "Host": req.host = value
			case "User-Agent": req.userAgent = value
			case "Accept": req.accept = strings.Split(value, ",")
			case "Authorization": req.authorization = value
		}
	}
	return req
}

