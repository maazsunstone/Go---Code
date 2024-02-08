// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func LanguageName() string {
	return
}

func main() {
	fmt.Println("Hello World!")
}
