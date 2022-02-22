package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.Args is slice
	for i, v := range os.Args {
		fmt.Println("index", i, "value ", v)
	}

	strings.Join(os.Args, " ")
}
