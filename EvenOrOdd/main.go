package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, " is Even")
		}else {
			fmt.Println(num, " is Odd")
		}
	}
}