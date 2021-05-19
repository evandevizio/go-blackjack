package main

import (
	d "deck/deck"
	"fmt"
)

var playerHand []d.Card
var dealerHand []d.Card

func main() {
	titleArt :=
		`
	========================================================================
	.------..------..------..------..------..------..------..------..------.
	|B.--. ||L.--. ||A.--. ||C.--. ||K.--. ||J.--. ||A.--. ||C.--. ||K.--. |
	|  ()  ||  /\  || (\/) ||  /\  ||  /\  ||  ()  || (\/) ||  /\  ||  /\  |
	| ()() || (__) ||  \/  ||  \/  || (__) || ()() ||  \/  ||  \/  ||  \/  |
	| '--'B|| '--'L|| '--'A|| '--'C|| '--'K|| '--'J|| '--'A|| '--'C|| '--'K|
	'------''------''------''------''------''------''------''------''------'
	========================================================================
		`
	fmt.Println(titleArt)
	deck := d.NewDeck()
	deck = d.ShuffleDeck(deck)
	//deal cards
	//main loop
}
