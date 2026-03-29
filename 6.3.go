package main

import (
	"fmt"
	"strings"
	"unicode"
)

func createUniqueText(text string) string {
	ret := ""
	valueExist := false
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	wordsMap := make(map[int]string)
	wordsSlice := strings.FieldsFunc(text, f)
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
		if len(wordsMap)-i == 1 {
			ret = ret + wordsMap[i]
		}
		ret = ret + wordsMap[i] + " "
	}
	return ret
}

func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	fmt.Println(text)
	fmt.Println(createUniqueText(text))
}
