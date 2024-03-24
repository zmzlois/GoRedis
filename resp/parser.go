package resp

import (
	"fmt"
	"strings"
)

type Resp struct{}

func (p *Resp) EncodeArray(input []string) ([]byte, error) {
	var result []byte

	formatStart := []byte(fmt.Sprintf("*%d", len(input)))
	result = formatStart

	for i := 0; i <= len(input)+1; i++ {

		if i <= len(input)-1 {
			char := input[i]

			stringFormat := fmt.Sprintf("%d\\r\\n%s", len(char), char)
			result = append(result, []byte("$")...)
			result = append(result, []byte(stringFormat)...)
		}

		if i == len(input)+1 {
			result = append(result, []byte("\\r\\n")...)
		}

	}

	return result, nil
}

func (p *Resp) Decode(input []byte) ([]string, error) {
	var result []string
	// Define regex patterns

	formatted := string(input)

	lines := strings.Split(formatted, "\r\n")

	for _, line := range lines {
		if line[0] != '*' && line[0] != '$' {
			result = append(result, line)
		}
	}

	return result, nil
}
