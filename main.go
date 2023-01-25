package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	//colors["white"] = "#FFFFFF"

	//delete(colors, "white")

	//printMap(colors)

	for color, hex := range colors{
		fmt.Println(color,"hex color is", hex)
	}
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color,"hex color is", hex)
	}
}