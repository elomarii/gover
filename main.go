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
		go handleConnection(conn)
	}
}

// fach tzid tkbr 4atwli module bo7do
func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Connection established with %s\n", clientAddr)

	// 4a brikol ta t7ydha
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 13\r\n" +
		"\r\n" +
		"Hello, World!"
	conn.Write([]byte(response))

	fmt.Printf("Response sent to %s\n", clientAddr)
}
