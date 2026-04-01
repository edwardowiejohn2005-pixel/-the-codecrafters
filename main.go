package main

import (
	"fmt"
	"strings"
)

func FixQuotes(words []string) []string {
	for i := 0; i < len(words); i++ {

		if words[i] == "'" && i+1 < len(words) {
			words[i+1] = "'" + words[i+1]
			words = append(words[:i], words[i+1:]...)
		}

		if i < len(words) && words[i] == "'" {
			words[i-1] = words[i-1] + "'"
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return words
}

func main() {
	input := "' I am the most well-known homosexual in the world '"
	words := strings.Fields(input)
	fmt.Println(FixQuotes(words))
}
