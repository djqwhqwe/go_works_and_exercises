package main

import (
	"fmt"
)

var Operate = func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i)
}

var Concat = func(xs ...interface{}) interface{} {
	ret := ""
	for _, value := range xs {
		if str, ok := value.(string); ok {
			ret += str
		}
	}
	return ret
}

var Sum = func(xs ...interface{}) interface{} {
	ret := 0.0
	for _, value := range xs {
		if v, ok := value.(int); ok {
			ret += v
		}
		if v, ok := value.(float64); ok {
			ret += v
		}
	}
	return ret
}

func main() {
	fmt.Println(Operate(Concat, "Hello, ", "World!"))  // Вывод: "Hello, World!"
	fmt.Println(Operate(Sum, 1, 2, 3, 4, 5))           // Вывод: 15
	fmt.Println(Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5)) // Вывод: 16.5
}
