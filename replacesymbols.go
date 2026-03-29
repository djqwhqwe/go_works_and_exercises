package main

import (
	"fmt"
	"strings"
)


func ReplaceSymbols(str string, old rune, new rune) string {
	return strings.ReplaceAll(str, string(old), string(new))
}

func main() {
	result := ReplaceSymbols("Hello, world!", 'o', '0')
	fmt.Println(result) // Выведет: Hell0, w0rld!

	//Дополнительные примеры
	result2 := ReplaceSymbols("Mississippi", 's', 'x')
	fmt.Println(result2) //Выведет: Mixxixxippi

	result3 := ReplaceSymbols("abcde", 'x', 'y')
	fmt.Println(result3) //Выведет: abcde (x не найдено)

}
