package deck

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var deck []Card

type Card struct {
	number int // The number of the card in sequence (A=1, J=11, Q=12, etc.)
	suit   string
	symbol string // A, 2, 3, ... , J, Q, K
}

func NewDeck() []Card {
	// Create a new factory default deck of 52 cards in order, no jokers.
	var newDeck []Card
	newDeck = make([]Card, 52)
	fmt.Println("Creating a new 52 card deck..." + "\n")

	for s := 0; s <= 12; s++ {
		newDeck[s].number = s + 1
		newDeck[s].suit = "Spades"
	}
	for d := 13; d <= 26; d++ {
		newDeck[d].number = (d + 1) - 13
		newDeck[d].suit = "Diamonds"
	}
	for c := 26; c <= 39; c++ {
		newDeck[c].number = (c + 1) - 26
		newDeck[c].suit = "Clubs"
	}
	for h := 39; h <= 51; h++ {
		newDeck[h].number = (h + 1) - 39
		newDeck[h].suit = "Hearts"
	}
	time.Sleep(2 * time.Second)
	FormatDeck(&newDeck)
	return newDeck
}

func FormatDeck(deck *[]Card) { // Delegate symbols
	for i := 0; i < len(*deck); i++ {
		if ((*deck)[i].number != 1) && ((*deck)[i].number != 11) &&
			((*deck)[i].number != 12) && ((*deck)[i].number != 13) {
			(*deck)[i].symbol = strconv.Itoa((*deck)[i].number)
		} else { // This deals with face cards
			switch (*deck)[i].number {
			case 1:
				(*deck)[i].symbol = "A"
			case 11:
				(*deck)[i].symbol = "J"
			case 12:
				(*deck)[i].symbol = "Q"
			case 13:
				(*deck)[i].symbol = "K"
			default:
				(*deck)[i].symbol = strconv.Itoa((*deck)[i].number)
			}
		}
	}
}

func PrintDeck(deck []Card) {
	cardMap := make(map[int]string)
	cardMap[1] = "A"
	cardMap[11] = "J"
	cardMap[12] = "Q"
	cardMap[13] = "K"
	suitMap := make(map[string]rune) // Unicode suit characters
	suitMap["Spades"] = '\u2660'
	suitMap["Diamonds"] = '\u2666'
	suitMap["Clubs"] = '\u2663'
	suitMap["Hearts"] = '\u2665'

	for i := 0; i < len(deck); i++ { // Display the cards
		fmt.Printf("%v"+" "+"%q"+"\n",
			deck[i].symbol, suitMap[deck[i].suit])
	}
	fmt.Print("\n")
}

func ShuffleDeck(deck []Card) []Card {
	shuffled := make([]Card, len(deck))
	fmt.Println("Shuffling the deck..." + "\n")

	r := rand.New(rand.NewSource(time.Now().Unix())) // Seed
	for i, randIndex := range r.Perm(len(deck)) {
		shuffled[i] = deck[randIndex] // Randomize the elements of deck
	}
	time.Sleep(2 * time.Second)
	return shuffled
}
