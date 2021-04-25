package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\033[31mbasename\033[0m: missing operand")
		os.Exit(1)
	}

	res := strings.Split(os.Args[1], "/")
	fmt.Println(res[len(res) - 1])
}
