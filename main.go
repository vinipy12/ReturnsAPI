package main

import "fmt"

func main() {
	server := newApiServer(":8080")
	fmt.Println("Starting Server on port 8080")
	server.Run()
}
