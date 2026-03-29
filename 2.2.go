package main

import (
	"testing"
)

func MaxDifference(numbers []int) int {
	if len(numbers) == 0 || len(numbers) == 1 {
		return 0
	}
	max := numbers[0]
	min := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return max - min
}

func TestMaxDifference(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1}
	c := []int{}
	if MaxDifference(a) != 3 || MaxDifference(b) != 0 || MaxDifference(c) != 0 {
		t.Errorf("Test has failed")
	}
}

func main() {}

