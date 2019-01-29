package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ? byte slice represents a string
// in a computer friendly way
// use type conversion
// []byte("Hi There!")

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// 0666 anyone can read and write file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
