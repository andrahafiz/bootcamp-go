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

	fmt.Println("Bentuk dasar", users)

	//Encode to JSON
	userJSON, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bentuk JSON", string(userJSON))

	// DECODE JSON KE OBJECT
	var users2 []User
	fmt.Println("User saat ini:", len(users2))

	// Data yang dikirim dari frontend
	dataDariFrontEnd := `{"Name":"NooBee"}`
	// Mengubah string menjadi byte
	data := []byte(dataDariFrontEnd)
	var user User

	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Tambah data gagal! \n", err.Error())
	}

	users2 = append(users2, user)
	fmt.Println("User stelah ditambah dari frontend", len(users))
	fmt.Println(users2)
	// fmt.Println(users2)


}
