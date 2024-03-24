package commands

import "fmt"

func Echo(stringArray []string) []byte {
	format := fmt.Sprintf("$%d\r\n%s\r\n", len(stringArray[1]), stringArray[1])

	result := []byte(format)

	return result
}
