package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("Ini adalah\nlangkah awal\nmenuju\nperjalanan baru")
	reader := bufio.NewReader(input)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Halo\n")
	_, _ = writer.WriteString("World\n")
	writer.Flush()

}
