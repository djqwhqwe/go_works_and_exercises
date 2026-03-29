package main

import (
	"fmt"
	"strings"
	"unicode"
)

func filterSentence(sentence string, filter map[string]bool) string {
	ret := ""
	valueExist := false
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	wordsMap := make(map[int]string)
	wordsSlice := strings.FieldsFunc(sentence, f)
	for i, jvalue := range wordsSlice {
		valueExist = false
		for _, zvalue := range wordsMap {
			if jvalue == zvalue {
				valueExist = true
			}
		}
		if valueExist == false {
			wordsMap[i] = jvalue
		}
	}
	for i := 0; i < len(wordsMap); i++ {
		if filter[wordsMap[i]] == true {
			continue
		}
		if len(wordsMap)-i == 1 {
			ret = ret + wordsMap[i]
		}
		ret = ret + wordsMap[i] + " "
	}
	return ret
}

func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}

	filteredSentence := filterSentence(sentence, filter)
	fmt.Println(filteredSentence)
}
