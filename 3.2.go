// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"testing"
)

func appendInt(xs *[]int, x ...int) {
	if len(x) > 0 {
		for i := range x {
			xs = append(xs, x[i])
		}
	}
}

func main() {

}

func TestAppedInt(t *testing.T) {
	a := []int{}
	a = appendInt(&a, 1, 2, 3))
	if 
}
