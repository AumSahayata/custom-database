package main

import (
	"encoding/json"
	"fmt"
)

const Version = "1.0.1"

type Address struct {
	City string
	State string
	Country string
	Pincode json.Number
}

type User struct {
	Name string
	Age json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

	employees := []User{
		{"John", "23", "0000000000", "Mangazon", Address{"Bangalore", "Karnataka", "India", "410013"}},
		{"Mango", "23", "1110000000", "ZTL", Address{"Bangalore", "Karnataka", "India", "410013"}},
		{"James", "19", "0000000432", "Kapple", Address{"Ahmedabad", "Gujarat", "India", "380060"}},
		{"Panter", "21", "1120450000", "Mangazon", Address{"Bangalore", "Karnataka", "India", "410013"}},
		{"Danto", "24", "0000000346", "FAANG", Address{"Pune", "Maharashtra", "India", "300013"}},
		{"Poket", "33", "0000000333", "Mangazon", Address{"Bangalore", "Karnataka", "India", "410013"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name: value.Name,
			Age: value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(records)

	allusers := []User

	for _, f := records {
		emloyeeFound := User{}
		if err := json.Unmarshal([]byte(f), &emloyeeFound); err != nil {
			fmt.Println("Error:", err)
		}
		allusers = append(allusers, emloyeeFound)
	}
	fmt.Println(allusers)

	// if err := db.Delete("user", "john"); err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// if err := db.Delete("user", ""); err != nil {
	// 	fmt.Println("Error:", err)
	// }


}