package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
			cards = append(cards, value+" of "+suit)
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

func newDeckFromFile(filename string) deck {
	// bs for byte slice
	// err object
	// if something went wrong then err will be populated
	// otherwise all good value with be nill
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck() to still make sure return a deck

		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// go random number generator
		// always uses the same seed value
		newPosition := r.Intn(len(d) - 1)

		// swap in one line
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
