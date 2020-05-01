package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type portugueseBot struct{}

func main() {
	eb := englishBot{}
	pb := portugueseBot{}

	printGreeting(eb)
	printGreeting(pb)
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

func (portugueseBot) getGreeting() string {
	// VERY custom logic for generating an portuguese greeting
	return "Olá!"
}

// func printGreeting(pb portugueseBot) {
// 	fmt.Println(pb.getGreeting())
// }
