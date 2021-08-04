package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := newCard() //"Ace of Spades"
	cards := deck{newCard(), "Three of Diamonds"}
	fmt.Println(card)
	fmt.Println(cards)
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
