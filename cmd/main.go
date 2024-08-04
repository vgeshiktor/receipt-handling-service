package main

import (
	"fmt"

	"github.com/vgeshiktor/receipt-handling-service/pkg/webserver"
)

func main() {
	server, err := webserver.NewHTTPServer("0.0.0.0", 8080)
	if err != nil {
		fmt.Printf("Error creating server: %v\n", err)
		return
	}

	server.Start()

}
