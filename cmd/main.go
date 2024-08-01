package main

import (
	"fmt"

	"github.com/vgeshiktor/rest_http_server/pkg/http"
)

func main() {
	server, err := http.NewHTTPServer("0.0.0.0", 8080)
	if err != nil {
		fmt.Printf("Error creating server: %v\n", err)
		return
	}

	server.Start()

}
