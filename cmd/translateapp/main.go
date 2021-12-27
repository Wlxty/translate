package main

import (
	"fmt"
	"os"
	"translateapp/internal/server"
)

func main() {
	// If you want to start server. Use Run method from server package.
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
