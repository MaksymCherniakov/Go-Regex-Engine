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

	if strings.Contains(regexp, "^") && strings.Contains(regexp, "$") {

		regexpValue := regexp[1:]
		regexpValue = regexpValue[:len(regexpValue)-1]

		if regexpValue == intput {
			return true
		}

	} else if strings.Contains(regexp, "^") {

		regexpValue := regexp[1:]

		if regexpValue == "." {
			return true
		} else if regexp[1:] == intput[:len(regexp[1:])] {
			return true
		}

	} else if strings.Contains(regexp, "$") {

		regexpValue := regexp[:len(regexp)-1]

		if regexpValue == "." {
			return true
		} else if regexp[:len(regexp)-1] == intput[len(regexp):] {
			return true
		}

	}

	if regexp == intput {
		return true
	}

	countPoint := strings.Count(regexp, ".")

	if countPoint == 1 && intput != "" {

		if regexp[0] == '.' {
			if strings.Contains(intput, regexp[1:]) {
				return true
			}
		}

		if regexp[len(regexp)-1] == '.' {
			if strings.Contains(intput, regexp[:len(regexp)-1]) {
				return true
			}
		}

	}

	if countPoint > 1 && len(regexp) == countPoint && intput != "" {
		return true
	}

	if countPoint == 0 && regexp != "" && intput != "" {
		if strings.Contains(intput, regexp) {
			return true
		}
	}

	if regexp == "" && intput != "" {
		return true
	}

	if regexp == "" && intput == "" {
		return true
	}

	return false

}
