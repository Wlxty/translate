package main

import (
	"fmt"
	"os"
	"translateapp/internal/translateapp"
)

func main() {
	//graceful shutdown to do
	if err := translateapp.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
