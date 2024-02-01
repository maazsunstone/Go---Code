package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")

	var bookingList = []string{"Maaz Abdullah", "Omar ibn Al Khattab", "Abu Bakr"}
	fmt.Println(bookingList)

	firstName := []string{}
	for _, booking := range bookingList {
		var names = strings.Fields(booking)
		firstName = append(firstName, names[0])
	}
	fmt.Println(firstName)
}
