package main

import (
	"fmt"

	"github.com/pborman/uuid"
)

var version string

func init() {
	version = "1"
}

func main() {
	uuid := uuid.NewRandom()
	fmt.Printf("version: %s, uuid: %s,content: hello world\n", version, uuid.String())
}
