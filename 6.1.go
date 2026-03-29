package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countWordOccurrences(text string) map[string]int {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	wordsMap := make(map[string]int)
	wordsSlice := strings.FieldsFunc(text, f)
	for _, value := range wordsSlice {
		wordsMap[value]++
	}
	return wordsMap
}

func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurrences := countWordOccurrences(text)
	for word, count := range occurrences {
		fmt.Printf("%s: %d\n", word, count)
	}
}
