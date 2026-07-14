package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	createNewFile("sample.log", "smaple log")
	addToFile("sample.log", " message string")
	result, _ := readFile("sample.log")
	fmt.Println(result)
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		message += string(line)
	}
	return message, nil
}

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}
