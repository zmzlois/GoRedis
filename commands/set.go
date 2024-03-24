package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Set(stringArray []string) []byte {
	var result []byte

	newLine := fmt.Sprintf("%s|%s", stringArray[1], stringArray[2])

	err := os.WriteFile("../store.txt", []byte(newLine), os.ModeAppend)

	Error(err)

	valueSet := ValidateSet(stringArray)

	if valueSet {
		result = []byte("+OK\r\n")
	}

	return result
}

func ValidateSet(stringArray []string) bool {
	result := false

	line := fmt.Sprintf("%s|%s", stringArray[1], stringArray[2])

	file, err := os.Open("../store.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == line {
			return true // Line found, validation successful
		}
	}

	if err := scanner.Err(); err != nil {
		Error(err)
	}

	return result
}
