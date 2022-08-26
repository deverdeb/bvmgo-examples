package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	serverAddress := ""
	if len(os.Args) != 2 {
		log.Printf(`You can use command arguments to specify server port.

Usage:
        %s [port]

The arguments are:
        port          server port - ex: 8562

Examples:
%s 8080

`, os.Args[0], os.Args[0])
	} else {
		serverAddress = ":" + os.Args[1]
	}

	fmt.Println("start server...")
	// listen client connection
	ln, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed to create TCP connection listener: %v", err)
	}
	defer ln.Close()
	fmt.Printf("server listen on %s (ctrl+c to exit)\n", ln.Addr().String())
	fmt.Println()

	// accept connection
	conn, err := ln.Accept()
	if err != nil {
		log.Fatalf("failed to accept connection: %v", err)
	}
	defer conn.Close()
	fmt.Printf("client %s connection\n", conn.RemoteAddr().String())

	// run loop forever (or until ctrl-c)
	for {
		// get message, output
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			fmt.Printf("client %s disconnection\n", conn.RemoteAddr().String())
			return
		}
		if err != nil {
			log.Fatalf("failed to receive message from client %s: %v", conn.RemoteAddr().String(), err)
		}
		message = strings.TrimSpace(message)
		fmt.Printf("< message received: %s\n", message)

		// send response
		fmt.Printf("> send response\n")
		_, err = fmt.Fprintf(conn, "hello client %s\n", conn.RemoteAddr().String())
		if err != nil {
			log.Fatalf("failed to send message to client %s: %v", conn.RemoteAddr().String(), err)
		}
	}
}
