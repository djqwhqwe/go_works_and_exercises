// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"testing"
)

func appendInt(xs []int, x ...int) []int {
	if len(x) > 0 {
		for i := range x {
			xs = append(xs, x[i])
		}
	}
	return xs
}

func main() {
	a := []int{}
	fmt.Println(appendInt(a, 1, 2, 3))
}

func TestAppedInt(t *testing.T) {
	a := []int{}
	fmt.Println(appendInt(a, 1, 2, 3))
}
