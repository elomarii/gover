package gover

// TODO parse HTTP request and
// -> identify HTTP version
// -> requested resource and method
// -> parse headers into a custom struct

type HTTPReq struct {
	method, resource, version string
	host                      string
	// ...
}

func parseHTTPRequest(request string) {

}
