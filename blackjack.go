package main

import (
	d "deck/deck"
	"fmt"
	"time"
)

var playerHand []d.Card
var dealerHand []d.Card

func Count(hand []d.Card) int {
	count := 0
	aceValue := 1
	for i := 0; i < len(hand); i++ {
		thisCard := hand[i]
		if thisCard.Number > 10 { // Face cards
			count += 10
		} else if thisCard.Number == 1 {
			fmt.Println("You have an ace. Treat it as 1 or 11? ")
			fmt.Scanln(&aceValue)
			switch aceValue {
			case 1:
				aceValue = 1
				count += aceValue
			case 11:
				aceValue = 11
				count += aceValue
			default:
				fmt.Println("Must enter 1 or 11.")
				panic(fmt.Sprintf("%v", i))
			}
		} else {
			count += hand[i].Number
		}
	}
	return count
}

func NaturalCheck(player []d.Card, dealer []d.Card) {
	if (player[0].Number == 1 || player[1].Number == 1) &&
		(player[0].Number >= 10 || player[1].Number >= 10) {
		fmt.Println("Blackjack!")
	} else if (dealer[0].Number == 1 || dealer[1].Number == 1) &&
		(dealer[0].Number >= 10 || dealer[1].Number >= 10) {
		fmt.Println("Blackjack!")
	}
}

func StartingHand(deck *[]d.Card, player *[]d.Card, dealer *[]d.Card) {
	fmt.Println("Dealing the starting hands..." + "\n")
	time.Sleep(2 * time.Second)
	playerHand = append(playerHand, d.DrawCard(deck, len(*deck)))
	dealerHand = append(dealerHand, d.DrawCard(deck, len(*deck)))
	playerHand = append(playerHand, d.DrawCard(deck, len(*deck)))
	dealerHand = append(dealerHand, d.DrawCard(deck, len(*deck)))
	NaturalCheck(playerHand, dealerHand)
	fmt.Println("Your hand: ")
	d.PrintDeck(playerHand, false)
	fmt.Println("Your score: ", Count(playerHand))
	fmt.Println("Dealer's hand: ")
	d.PrintDeck(dealerHand, true)
}

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

	StartingHand(&deck, &playerHand, &dealerHand)

	//main loop
}
