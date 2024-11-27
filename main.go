package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	port := ":80"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Server is listening on port %s\n", port)

	// main loop, listen for connections and handle requests
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		clientAddr := conn.RemoteAddr().String()

		requestBuffer := make([]byte, 1024)
		n, err := conn.Read(requestBuffer)
		if err != nil {
			fmt.Println(clientAddr, "NONE", "NONE", 500)
		}
		req := string(requestBuffer[:n])
		request := ParseHTTPRequest(req)
		
		var statusCode int

		switch request.method {
			case "GET": statusCode = httpGet(conn, request.resource) 
		}

		fmt.Println(clientAddr, request.method, request.resource, statusCode)
	}
}

func httpGet(conn net.Conn, path string) int {
	defer conn.Close()

	content, err := os.ReadFile("./" + path)
	if err != nil {
		return 404
	}
	header := "HTTP/2 200 OK\r\n\n" +
		string(content)
	conn.Write([]byte(header))

	return 200
}
