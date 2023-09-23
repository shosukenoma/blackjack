package main

import (
	"fmt"
)

type card struct{
	suit string
	rank string
	value int
}
type deck []card

func newDeck(includeJoker bool) deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardRanks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, suit := range cardSuits {
		for index, rank := range cardRanks {
			card := card{suit, rank, index + 1}
			cards = append(cards, card)
		}
	}

	// if includeJoker {
	// 	cards = append(cards, "Z", "Z")
	// }
	return cards
}

func deal(d deck, handSize int) {
	myHand := d[:handSize]
	opponentsHand := d[handSize: handSize*2]
	remainingDeck := d[handSize*2:]

	fmt.Println(myHand)
	fmt.Println(opponentsHand)
	fmt.Println(remainingDeck)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}