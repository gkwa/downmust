package main

import (
	"fmt"
	"os"

	"github.com/taylormonacelli/downmust/version"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-version" || os.Args[1] == "version") {
		fmt.Println(version.GetBuildInfo())
		os.Exit(0)
	}

	fmt.Println("Hello, World!")
}
