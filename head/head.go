package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func print_lines(file *os.File, n uint64) {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var i uint64 = 0
	for ; i < n && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
}

func main() {

	// POSIX mandates that the last -n flag is the one that is used for all files
	var n uint64 = 10

	// If there is more than one file to print, you need to add the filenames
	var filenames []string
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-n" {
			if i + 1 >= len(os.Args) {
				fmt.Println("\033[31mhead\033[0m: expected integer after '-n' flag")
				os.Exit(1)
			}
			res, err := strconv.ParseUint(os.Args[i + 1], 10, 64)
			if err != nil {
				fmt.Println("\033[31mhead\033[0m: invalid parameter to '-n' flag")
				os.Exit(1)
			}
			n = res
			i++
		} else {
			filenames = append(filenames, os.Args[i])
		}
	}

	if len(filenames) == 0 {
		print_lines(os.Stdin, n)
	} else if len(filenames) == 0 {
		file, err := os.Open(filenames[0])
		if err != nil {
			fmt.Printf("\033[31mhead\033[0m: cannot open '%s'\n", filenames[0])
			os.Exit(1)
		}
		print_lines(file, n)
	} else {
		for i, filename := range filenames {
			if filename == "-" {
				if i == 0 {
					fmt.Printf("===> standard in  <===\n")
				} else {
					fmt.Printf("\n===> standard in <===\n")
				}
				print_lines(os.Stdin, n)
			} else {
				file, err := os.Open(filename)
				if err != nil {
					fmt.Printf("\033[31mhead\033[0m: cannot open '%s'\n", filename)
					os.Exit(1)
				}
				if i == 0 {
					fmt.Printf("===> %s <===\n", filename)
				} else {
					fmt.Printf("\n===> %s <===\n", filename)
				}
				print_lines(file, n)
			}
		}
	}
}
