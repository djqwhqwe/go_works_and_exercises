// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"

	"github.com/brianvoe/gofakeit"
)

type User struct {
	Name string
	Age  int
}

func getUsers() []User {
	slice := []User{}
	gofakeit.Seed(int64(gofakeit.NanoSecond()))
	for i := 0; i < 10; i++ {
		slice = append(slice, User{Name: gofakeit.Name(), Age: gofakeit.Number(18, 60)})
	}
	return slice
}

func preparePrint(users []User) string {
	info := ""
	for i := 0; i < 10; i++ {
		info += "Имя: " + users[i].Name + ", " + "Возраст: " + strconv.Itoa(users[i].Age) + "\n"
	}
	return info
}

func main() {
	users := getUsers() // Получаем срез 10 пользователей
	result := preparePrint(users)
	fmt.Println(result)
}
