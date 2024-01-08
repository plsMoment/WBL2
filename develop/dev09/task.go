package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var StatusCodeError = errors.New("status differ from 200")

// wget realize the simplest example of wget utility.
// Downloads from url HTML page, non-recursive
func wget(url string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	lastSlashIdx := strings.LastIndex(url, "/")
	var filename string
	if lastSlashIdx != -1 {
		filename = url[lastSlashIdx:]
	} else {
		filename = "/index.html"
	}
	file, err := os.Create(dir + filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return StatusCodeError
	}

	_, err = io.Copy(file, resp.Body)

	return err
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("too many args, enter only url")
	}
	url := os.Args[1]

	if err := wget(url); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success")
}
