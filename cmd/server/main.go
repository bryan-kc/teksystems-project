package main

import (
	"fmt"
	"github.com/bryan-kc/teksystems-project/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
