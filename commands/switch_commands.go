package commands

import (
	"fmt"

	"github.com/zmzlois/GoRedis/resp"
)

type Store map[string]string

func SwitchCommands(buf []byte) []byte {
	var result []byte

	resp := resp.Resp{}

	stringArray, err := resp.Decode(buf)
	if err != nil {
		fmt.Println("handle.connection.decode.error:", err)
	}

	switch command := stringArray[0]; command {

	case "echo":
		result = Echo(stringArray)

	case "set":
		result = Set(stringArray)

	case "get":
		result = Get(stringArray)

	default:
		result = []byte("+PONG\r\n")
	}

	return result
}
