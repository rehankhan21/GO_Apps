package main

import "fmt"

func main() {
	// creating a map where all keys are type string and all values are type string as well
	colors := map[string]string{
		"red":   "#ff0000",
		"greem": "#4bf745",
	}

	fmt.Println(colors)
}