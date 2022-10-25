package main

import (
	"fmt"
	"os"

	"github.tools.sap/I521862/circuit-breaker-demo/client/internal"
)

func main() {
	server := internal.CreateServer()

	if os.Getenv("PORT") == "" {
		server.Logger.Fatal(server.Start(":8080"))
	} else {
		server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
	}
}
