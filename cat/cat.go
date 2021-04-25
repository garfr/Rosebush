package main

import (
	"fmt"
	"flag"
	"io"
	"bufio"
	"os"
)

func main() {
	// -u seems to be ignored by gnu and busybox, so we'll ignore it also
	flag.Bool("u", false, "Use unbuffered IO (ignored)")

	flag.Parse()
	args := flag.Args()
	reader := bufio.NewReader(os.Stdin)

	if len(args) == 0 {
		readStdin(reader)
		return
	}

	for i := range args {
		file := args[i]
		if file == "-" {
			readStdin(reader)
			continue
		}

		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}
		os.Stdout.Write(data)
	}
}

func readStdin(reader *bufio.Reader) {
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		os.Stdout.Write(line)
	}
}
