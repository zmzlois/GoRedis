package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/zmzlois/GoRedis/resp"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Start handling connection...")
	defer conn.Close()

	buf := make([]byte, 4086)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("Error: ", err)
			break
		}

		if err != nil {
			fmt.Println("Error reading from client:", err.Error())
			os.Exit(1)
		}

		log.Println("Read", n)

		result, err := handleEcho(buf)
		if err != nil {
			fmt.Println("handleConnection.error", err)
		}

		_, err = conn.Write(result)
		if err != nil {
			return
		}
	}
}

func handleEcho(n []byte) ([]byte, error) {
	var result []byte

	resp := resp.Resp{}

	stringArray, err := resp.Decode(n)
	if err != nil {
		fmt.Println("handleEcho.error", err)
	}

	switch command := stringArray[0]; command {
	case "echo":

		format := fmt.Sprintf("$%d\r\n%s\r\n", len(stringArray[1]), stringArray[1])

		result = []byte(format)

	default:
		result = []byte("+PONG\r\n")
	}

	return result, nil
}
