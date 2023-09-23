package main

func main() {
	cards := newDeck(false)
	cards.shuffle()
	// deal(cards, 2)
	cards.print()
}