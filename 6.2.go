package main

import (
	"fmt"
	"maps"
)

func mergeMaps(map1, map2 map[string]int) map[string]int {
	maps.Copy(map1, map2)
	return map1
}

func main() {

	map1 := map[string]int{"apple": 3, "banana": 2}
	map2 := map[string]int{"orange": 5, "grape": 4}
	for key, value := range mergeMaps(map1, map2) {
		fmt.Printf("%s: %d\n", key, value)
	}
}