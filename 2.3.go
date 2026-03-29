package main

import (
	"fmt"
	"testing"
)


func bitwiseXOR(n, res int) int {
	return res ^ n
}

func findSingleNumber(numbers []int) int {
	result := 0
	index:=0
	for i, num := range numbers {
		result = bitwiseXOR(num, result)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	singleNumber := findSingleNumber(numbers)
	fmt.Println(singleNumber) // 5
}

func test(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	if findSingleNumber(numbers) != 5{
		t.Errorf("Fuck")
	}
}

