package main

import (
	"fmt"
	"regexp"
)

func main() {
	patern := `^[a-zA-z]+`
	re := regexp.MustCompile(patern)
	result := re.MatchString("aku cinta dan sayang dia")
	fmt.Println(result)
}
