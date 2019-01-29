package main

import "fmt"

// create a new type of 'deck
// which is a slice of strings
// a deck which is equal to a slice of strings
type deck []string

// doesn't need a receiver
func newDeck() deck {
	cards := deck{}

	cardSuits := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardvalues := [4]string{"Ace", "Two", "Three", "Four"}

	// replace variable with underscore you don't care about
	for _, suit := range cardSuits {
		for _, value := range cardvalues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// receiver function
// receiver d of type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal function
// arguments inside function call
// return two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
