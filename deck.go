package main

import "fmt"

type deck []string

func newDeck(includeJoker bool) deck {
	cards := deck{}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for i := 0; i < 4; i++ {
		for _, value := range cardValues {
			cards = append(cards, value)
		}
	}
	if includeJoker {
		cards = append(cards, "Z", "Z")
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}