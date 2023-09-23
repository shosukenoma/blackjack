package main

func main() {
	cards := newDeck(false)
	cards.shuffle()
	myHand, opponentsHand, remainingDeck := deal(cards, 2)

	myHand.print()
	opponentsHand.print()
	remainingDeck.print()
}