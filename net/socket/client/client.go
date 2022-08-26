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
	var serverAddress string
	reader := bufio.NewReader(os.Stdin)

	if len(os.Args) != 2 {
		log.Printf(`You can use command arguments to specify server.

Usage:
        %s [server]

The arguments are:
        server        server address - ex: localhost:8562

Examples:
%s 192.168.1.12:52100

`, os.Args[0], os.Args[0])
		fmt.Print("enter server address (ex: 192.168.1.12:52100): ")
		serverAddress, _ = reader.ReadString('\n')
		fmt.Println()
	} else {
		serverAddress = os.Args[1]
	}

	serverAddress = strings.TrimSpace(serverAddress)

	fmt.Println("start client...")
	fmt.Println("connect to server", serverAddress)
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed to connect to server %s: %v", serverAddress, err)
	}
	defer conn.Close()

	fmt.Println("enter empty message or ctrl+c to stop")
	for {
		fmt.Println()
		fmt.Print("enter new message: ")
		messageToServer, _ := reader.ReadString('\n')
		messageToServer = strings.TrimSpace(messageToServer)
		if len(messageToServer) == 0 {
			return // empty message: stop client
		}

		// send message to server
		fmt.Println("> send message:", messageToServer)
		_, err = fmt.Fprintln(conn, messageToServer)
		if err != nil {
			log.Fatalf("failed to send message to server %s: %v", serverAddress, err)
		}

		// wait for reply
		messageFromServer, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			fmt.Printf("server %s disconnection\n", conn.RemoteAddr().String())
			return // server disconnection: exit
		} else if err != nil {
			log.Fatalf("failed to receive message from server %s: %v", serverAddress, err)
		}
		messageFromServer = strings.TrimSpace(messageFromServer)
		fmt.Printf("< receive response: %s\n", messageFromServer)
	}
}
