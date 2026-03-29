package main

import (
	"fmt"
)

var vowelMap map[rune]bool

func init() {
	vowelMap = make(map[rune]bool)
	for _, t := range []rune("aeiouyAEIOUYаеёиоуыэюяАЕЁИОУЫЭЮЯ") {
		vowelMap[t] = true
	}
}

func CountVowels(str string) int {
	count := 0
	for _, value := range []rune(str) {
		_, ok := vowelMap[value]
		if ok {
			count++
		}
	}
	return count
}

func main() {
	// Пример использования функции CountVowels

	count := CountVowels("Привет, мир!")
	fmt.Println(count) // Вывод: 3
	count = CountVowels("Hello, worldaaa!")
	fmt.Println(count) // Вывод: 3
}
