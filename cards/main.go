package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"   // declaring new variable with inferred type
	card = "Five of Diamonds" // reassigning value to it

	fmt.Println(card)
}
