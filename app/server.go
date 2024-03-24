package main

import (
	"fmt"
	"io"

	// Uncomment this block to pass the first stage
	"net"
	"os"

	"github.com/zmzlois/GoRedis/commands"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	// try?

	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		// does this handle concurrency?

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Start handling connection...")
	defer conn.Close()

	buf := make([]byte, 4086)

	for {
		_, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("Error: ", err)
			break
		}

		if err != nil {
			fmt.Println("Error reading from client:", err.Error())
			os.Exit(1)
		}

		result := commands.SwitchCommands(buf)

		_, err = conn.Write(result)
		if err != nil {
			return
		}
	}
}
