package main

import (
  "os"
  "strings"
  "fmt"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("\033[31mbasename\033[0m: missing operand\n")
    os.Exit(1)
  }

  res := strings.Split(os.Args[1], "/")
  fmt.Println(res[len(res)-1])
}
