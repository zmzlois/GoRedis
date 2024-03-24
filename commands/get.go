package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Get(stringArray []string) []byte {
	var result []byte

	getKey := stringArray[1]

	file, err := os.Open("../store.txt")

	Error(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := string(strings.TrimSpace(scanner.Text()))

		key := strings.Split(line, "|")[0]

		value := strings.Split(line, "|")[1]

		if key == getKey {
			result = []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(value), value))
		}

		if key != getKey {
			// if no key is found return "null bulk string"
			result = []byte("$-1\r\n")
		}
	}

	return result
}
