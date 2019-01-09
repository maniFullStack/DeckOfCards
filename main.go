package main

func main() {
	cards := newDeckFromFile("mycards")
	cards.shuffle()
	cards.print()

	//cards := newDeck()
	//cards.saveToFile("myCards")

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
}
