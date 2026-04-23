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

	if strings.HasPrefix(regexp, "^") {
		Println(matchSymbols(regexp[1:], intput))
	} else {
		Println(search(regexp, intput))
	}

}

func search(regexp, intput string) bool {

	if matchSymbols(regexp, intput) {
		return true
	}

	if intput != "" {
		return search(regexp, intput[1:])
	}
	return false
}

func matchSymbols(regexp, intput string) bool {

	if regexp == "" {
		return true
	}
	if regexp == "$" && intput == "" {
		return true
	}

	if intput == "" {
		return false
	}

	if len(regexp) > 1 {
		char := regexp[0]
		op := regexp[1]
		if op == '?' {
			return matchSymbols(regexp[2:], intput) ||
				(check(char, intput[0]) && matchSymbols(regexp[2:], intput[1:]))
		}
		if op == '*' {
			return matchSymbols(regexp[2:], intput) ||
				(check(char, intput[0]) && matchSymbols(regexp, intput[1:]))
		}
		if op == '+' {
			return check(char, intput[0]) &&
				(matchSymbols(regexp[2:], intput[1:]) || matchSymbols(regexp, intput[1:]))
		}
	}

	if check(regexp[0], intput[0]) {
		return matchSymbols(regexp[1:], intput[1:])
	}
	return false
}

func check(r, i byte) bool {
	return r == i || r == '.'
}
