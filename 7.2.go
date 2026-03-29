package main

import (
	"fmt"
)

// Функция getBytes возвращает срез байтов строки s
func getBytes(s string) []byte {
	return []byte(s)
}

// Функция getRunes возвращает срез символов строки s
func getRunes(s string) []rune {
	return []rune(s)
}

func main() {
	bytes := getBytes("Привет, мир!")
	fmt.Println(bytes) // Вывод: 21

	// Пример использования функции countSymbols
	symbols := getRunes("Привет, мир!")
	fmt.Println(symbols) // Вывод: 12
}
