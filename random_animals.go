// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"

	"github.com/brianvoe/gofakeit"
)

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	slice := []Animal{}
	gofakeit.Seed(int64(gofakeit.NanoSecond()))
	for i := 0; i < 10; i++ {
		slice = append(slice, Animal{Type: gofakeit.AnimalType(), Name: gofakeit.PetName(), Age: gofakeit.Number(0, 60)})
	}
	return slice
}

func preparePrint(animals []Animal) string {
	info := ""
	for i := 0; i < 10; i++ {
		info += "Тип животного: " + animals[i].Type + ", " + animals[i].Name + ", " + "Возраст: " + strconv.Itoa(animals[i].Age) + "\n"
	}
	return info
}

func main() {
	animals := getAnimals()
	result := preparePrint(animals)
	fmt.Println(result)
}
