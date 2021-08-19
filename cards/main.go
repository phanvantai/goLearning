package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//card := newCard() //"Ace of Spades"
	//cards := deck{newCard(), "Three of Diamonds"}
	cards := newDeck()

	fmt.Println(cards)

	cards.print()
}
