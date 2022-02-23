package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "copy failed: %v\n", err)
		os.Exit(1)
	}
}
