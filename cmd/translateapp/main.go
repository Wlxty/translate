package main

import (
	"fmt"
	"os"
	"translateapp/internal/server"
)

func main() {
	//graceful shutdown to do
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
