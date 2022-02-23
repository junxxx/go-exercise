package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/junxxx/go-exercise/1.1/dup"
)

func main() {
	begin := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("runtime : %d ns\n", time.Since(begin).Nanoseconds())

	start := time.Now()
	strings.Join(os.Args[1:], " ")
	fmt.Printf("lib runtime : %d ns\n", time.Since(start).Nanoseconds())
	dup.Dup()
}
