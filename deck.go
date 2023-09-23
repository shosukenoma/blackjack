package main

import (
	"fmt"
	"math/rand"
	"time"
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

func (d deck)shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[newPosition], d[i] = d[i], d[newPosition]
	}
}

func deal(d deck, handSize int) (deck, deck, deck){
	myHand := d[:handSize]
	opponentsHand := d[handSize: handSize*2]
	remainingDeck := d[handSize*2:]

	return myHand, opponentsHand, remainingDeck
}

func hit(hand deck, remaining deck) (deck, deck){
	hand = append(hand, remaining[0])
	remaining = remaining[1:]
	return hand, remaining
}

func (d deck) print() {
	c := []string{}
	for _, card := range d {
		c = append(c, "{", card.suit, card.rank, "}")
	}
	fmt.Println(c)
}