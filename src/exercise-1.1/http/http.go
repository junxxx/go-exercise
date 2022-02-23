package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Get(url string) {
	preFix := "http://"
	if !strings.HasPrefix(url, preFix) {
		url = preFix + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "copy failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nresponse status: %d\n", resp.StatusCode)
}
