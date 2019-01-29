package main

// array primitive fixed length list
// slice array that can grow or shrink
// must be defined with same type

// can only initialize variables outside of function
// cannot assign variables

func main() {
	// variable we referred to as card
	// ? var | name | type ever assigned
	// static typed
	// basic go types => bool, string, int, float64
	// var card string = "Ace of Spades" // one way of defining variable in go
	// var card string // just initializes the variable

	// ! := will infer types ONLY WHEN DEFINING NEW VARIABLE
	/*
		card := "Ace of Spades"   // relying on code compiler to figure out
		card = "Five of Diamonds" // only use colon on initialization
	*/

	// ? SLICE ──────────────────────────────────────────────────────────────────────
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	cards := newDeck()

	// iterate over slice or closed set
	// i is just the index
	// declare second variable as well to grab data
	// ? for index || declare card := range of cards
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
}

// ? func | name | type
// func newCard() string {
// 	return "Five of Diamonds"
// }
