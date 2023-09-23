package main

import (
	"fmt"
)

func main() {
	cards := newDeck(false)
	cards.shuffle()
	myHand, opponentsHand, remainingDeck := deal(cards, 2)

	myHand.print()
	opponentsHand.print()
	fmt.Println("Number of cards remaining in the deck:", len(remainingDeck))
}