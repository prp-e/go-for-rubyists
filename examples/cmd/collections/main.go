package main

import (
	"fmt"
	"slices"
	"unicode/utf8"
)

func main() {
	books := []string{"Dune", "Solaris", "Kindred"}
	first := books[:2]
	first[0] = "Neuromancer"
	fmt.Println(books[0])

	copyOfFirst := slices.Clone(first)
	copyOfFirst[0] = "Dune"
	fmt.Println(books[0])

	ratings := map[string]int{"Dune": 0}
	for _, title := range []string{"Dune", "Solaris"} {
		rating, found := ratings[title]
		fmt.Println(title, rating, found)
	}

	text := "café"
	fmt.Println(len(text), utf8.RuneCountInString(text))
}
