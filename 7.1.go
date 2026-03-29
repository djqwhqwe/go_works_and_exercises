package main

import (
	"fmt"
)

func countBytes(str string) int {
	return len(str)
}

func countSymbols(str string) int {
	return len([]rune(str))
}

func main() {
	bytes := countBytes("Привет, мир!")
	fmt.Println(bytes) // Вывод: 21

	// Пример использования функции countSymbols
	symbols := countSymbols("Привет, мир!")
	fmt.Println(symbols) // Вывод: 12
}