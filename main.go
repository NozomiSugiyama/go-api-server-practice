package main

import (
	"fmt"
	"github.com/NozomiSugiyama/go-api-server-practice/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}
