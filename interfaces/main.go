package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}
	printGreeting(eb)
	printGreeting(sp)
}

func (englishBot) getGreeting() string {
	// VERY custom logic to get the greeting
	return "Hi!"
}
func (spanishBot) getGreeting() string {
	// VERY custom logic to get the greeting
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
