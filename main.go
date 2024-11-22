package main

import (
	"fmt" // dakchi dial input output
	"net" // sockets rak fahm
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

		// Handle the connection in a separate goroutine
		clientAddr := conn.RemoteAddr().String()
		fmt.Printf("Connection established with %s\n", clientAddr)

		requestBuffer := make([]byte, 1024)
		n, err := conn.Read(requestBuffer)
		if err != nil {
			httpHandleError(err)
		}
		req := string(requestBuffer[:n])
		request := ParseHTTPRequest(req)
		fmt.Println(request.method)
		fmt.Println(request.host)
		//args := strings.Split(request, " ")
		//method := args[0]
		//switch method {
		//case "GET":
		//	httpGet(conn, args[1])
		//}
	}
}

func httpGet(conn net.Conn, path string) {
	defer conn.Close()

	content, err := os.ReadFile("./" + path)
	if err != nil {
		httpHandleError(err)
	}
	header := "HTTP/2 200 OK\r\n" +
		"Content-Type: text/plain\r\n\n" +
		string(content)
	conn.Write([]byte(header))
}

func httpHandleError(err error) {
	fmt.Println("error occured", err)
}
