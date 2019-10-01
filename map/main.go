package main

import "fmt"

func main() {
	// colors := make(map[string]string)

	// var colors map[string]string

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffffff" // must always use brackets for maps

	delete(colors, "white")

	printMap(colors)
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
