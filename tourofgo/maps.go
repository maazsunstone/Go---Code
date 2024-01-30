package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "I love Go!"
	words := strings.Fields(s)
	counts := make(map[string]int, len(words))
	fmt.Println(counts)
	for _, word := range words {
		counts[word]++
	}
	fmt.Println(counts)

}
