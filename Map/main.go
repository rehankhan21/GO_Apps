package main

import "fmt"

func main() {
	// creating a map where all keys are type string and all values are type string as well
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"greem": "#4bf745",
	// }

	// another way to make a map
	//var colors map[string]string

	colors := make(map[int]string)

	colors[10] = "#ffffff"
	
		delete(colors, 10)

	fmt.Println(colors)
}