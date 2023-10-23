package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Class string
}

func main() {
	users := []User{
		User{Name: "NooB", Class: "A"},
		User{Name: "Java", Class: "B"},
		User{Name: "Docker", Class: "C"},
	}

	data := saveToFile(users)

	fmt.Println("Data pengguna yang dibaca dari file JSON:")
	load := loadFromFile(data)
	for _, x := range load {
		fmt.Print("Nama: ", x.Name, ", ")
		fmt.Print("Class: ", x.Class, "\n")
	}
}

func saveToFile(user []User) string {
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err.Error())
		return "-"
	}
	fmt.Println("Data pengguna berhasil disimpan dalam file JSON")
	return string(userJSON)
}

func loadFromFile(data string) []User {

	var user []User
	var err = json.Unmarshal([]byte(data), &user)

	if err != nil {
		fmt.Println(err.Error())
		return user
	}
	return user
}
