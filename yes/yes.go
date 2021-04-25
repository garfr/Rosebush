package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := "y"

	if len(os.Args) > 1 {
		str = strings.Join(os.Args[1:], " ")
	}

	for {
		fmt.Println(str)
	}
}
