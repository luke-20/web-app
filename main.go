package main

import (
	"fmt"

	"github.com/luke-20/web-app/httpserver"
)

func main() {
	fmt.Println("Main!")

	// server := httpserver.Server{}
	// server.InitServer()
	// server.RunEndpoints()

	httpserver.Hlavni()
	// database.Hlav()
}
