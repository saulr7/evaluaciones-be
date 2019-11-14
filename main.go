package main

import (
	"fmt"

	"./server"
)

func main() {
	fmt.Println("Running...")
	server.Serve()
}
