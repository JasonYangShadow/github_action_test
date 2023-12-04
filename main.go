package main

import "fmt"

var version string

func init() {
	version = "1"
}

func main() {
	fmt.Printf("version: %s, content: hello world\n", version)
}
