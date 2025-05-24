package main

import (
	"fmt"
)

func main() {
	//мапа - це невпорядкована колекція пар ключ-значень, де ключі УНІКАЛЬНІ
	//var name map[KeyType]ValueType
	//var scores map[string]int

	userAges := map[string]int{
		"Oleg":      52,
		"Oleksandr": 49,
		"Jesse":     125,
	}
	fmt.Println(userAges)

	contacts := make(map[string]string)

	contacts["Oleksandr"] = "+3800505050123"
	contacts["Oleg"] = "+380123456789010"

	fmt.Println(contacts)

	phone, exists := contacts["Oleg"]
	if exists {
		fmt.Println(phone)
	} else {
		fmt.Println("nothing")
	}

	delete(contacts, "Oleg")
	fmt.Println(contacts)
}
