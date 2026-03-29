package main

import "fmt"

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	k:=0
	nicknameMap := make(map[User]User)
	for _, user := range users {
		_, ok := nicknameMap[user.Nickname]
		if !ok {
			nicknameMap[user.Nickname] = user
			k+=1
		}
	}
	fmt.Print(nicknameMap)
	return users
	//return uniqueUsers[:len(uniqueUsers):len(uniqueUsers)]
}

func main() {
	users := []User{
		{Nickname: "user1", Age: 30, Email: "user1@example.com"},
		{Nickname: "user1", Age: 35, Email: "user1_new@example.com"},
		{Nickname: "user2", Age: 25, Email: "user2@example.com"},
		{Nickname: "user3", Age: 40, Email: "user3@example.com"},
		{Nickname: "user2", Age: 28, Email: "user2_new@example.com"},
	}

	uniqueUsers := getUniqueUsers(users)
	uniqueUsers = append(	uniqueUsers, User{Nickname: "user1", Age: 30, Email: "user1@example.com"})
	//fmt.Println(uniqueUsers)
	//fmt.Println("Capacity:", cap(uniqueUsers)) // Demonstrates capacity
}
