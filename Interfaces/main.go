package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println((b.getGreeting()))
}

// if we are not using the value we can just have the reviever type alone
func (englishBot) getGreeting() string {
	// custom logic
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }