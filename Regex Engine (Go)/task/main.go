package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	inputValue := scanner.Text()

	parts := strings.Split(inputValue, "|")

	regexp := parts[0]
	intput := parts[1]

	result := matchSymbols(regexp, intput)
	Println(result)

}

func matchSymbols(regexp, intput string) bool {

	var result bool

	if regexp == intput {
		result = true
	} else if regexp == "." && intput != "" {
		result = true
	} else if regexp == "" && intput != "" {
		result = true
	} else if regexp == "" && intput == "" {
		result = true
	} else {
		result = false
	}

	return result

}
