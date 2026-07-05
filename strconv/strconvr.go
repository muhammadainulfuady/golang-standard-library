package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrParseBool = errors.New("Tidak bisa konversi ke boolean")
)

func parseBool(value string) (bool, error) {
	result, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("%w: %v", ErrParseBool, err)
	}
	return result, nil
}

func main() {
	result, err := parseBool("true")
	if err != nil {
		if errors.Is(err, ErrParseBool) {
			fmt.Println(err)
		}
	}
	fmt.Println(result)
}
