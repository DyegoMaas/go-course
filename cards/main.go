package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// hand, _ := deal(cards, 5)
	// hand.saveToFile("my_cards")

	// hand = newDeckFromFile("my_cards")
	// hand.print()
}
