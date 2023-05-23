package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clover"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (cards1 deck) print() {
	for i, card := range cards1 {
		fmt.Println(i, card)
	}
}
func (cards2 deck) print2() {
	for i, card := range cards2 {
		fmt.Println(i, card, "cards2")
	}
}
