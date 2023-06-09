package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	url := "https://github.com"
	protocol, domain := urlCutter(url)
	createAndWrite("test", protocol, domain)
}

func urlCutter(url string) (string, string) {
	protocol, domain, found := strings.Cut(url, "://")
	defer fmt.Println("found:", found)
	return "protocol: " + protocol, "domain: " + domain
}

func createAndWrite(filename string, text ...string) error {
	stringToWrite := strings.Join(text, "\n")
	filepath := "static/" + filename + ".txt"
	fmt.Println(filepath)

	var err error
	filePointer, err := os.Create(filepath)
	if err != nil {
		return err
	}
	_, err = io.WriteString(filePointer, stringToWrite)
	if err != nil {
		return err
	}
	defer filePointer.Close()
	return filePointer.Close()
}
