package commands

import (
	"fmt"
	"os"
)

func Error(e error) {
	cwd, _ := os.Getwd()

	if e != nil {
		fmt.Printf("%s error:%v", cwd, e)
	}
}
