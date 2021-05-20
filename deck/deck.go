package deck

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var deck []Card

type Card struct {
	Number int // The Number of the card in sequence (A=1, J=11, Q=12, etc.)
	Suit   string
	Symbol string // A, 2, 3, ... , J, Q, K
}

func NewDeck() []Card {
	// Create a new factory default deck of 52 cards in order, no jokers.
	var newDeck []Card
	newDeck = make([]Card, 52)
	fmt.Println("Creating a new 52 card deck..." + "\n")

	for s := 0; s <= 12; s++ {
		newDeck[s].Number = s + 1
		newDeck[s].Suit = "Spades"
	}
	for d := 13; d <= 26; d++ {
		newDeck[d].Number = (d + 1) - 13
		newDeck[d].Suit = "Diamonds"
	}
	for c := 26; c <= 39; c++ {
		newDeck[c].Number = (c + 1) - 26
		newDeck[c].Suit = "Clubs"
	}
	for h := 39; h <= 51; h++ {
		newDeck[h].Number = (h + 1) - 39
		newDeck[h].Suit = "Hearts"
	}
	time.Sleep(2 * time.Second)
	FormatDeck(&newDeck)
	return newDeck
}

func FormatDeck(deck *[]Card) { // Delegate Symbols
	for i := 0; i < len(*deck); i++ {
		if ((*deck)[i].Number != 1) && ((*deck)[i].Number != 11) &&
			((*deck)[i].Number != 12) && ((*deck)[i].Number != 13) {
			(*deck)[i].Symbol = strconv.Itoa((*deck)[i].Number)
		} else { // This deals with face cards
			switch (*deck)[i].Number {
			case 1:
				(*deck)[i].Symbol = "A"
			case 11:
				(*deck)[i].Symbol = "J"
			case 12:
				(*deck)[i].Symbol = "Q"
			case 13:
				(*deck)[i].Symbol = "K"
			default:
				(*deck)[i].Symbol = strconv.Itoa((*deck)[i].Number)
			}
		}
	}
}

func PrintDeck(deck []Card, dealerFlag bool) {
	CardMap := make(map[int]string)
	CardMap[1] = "A"
	CardMap[11] = "J"
	CardMap[12] = "Q"
	CardMap[13] = "K"
	SuitMap := make(map[string]rune) // Unicode Suit characters
	SuitMap["Spades"] = '\u2660'
	SuitMap["Diamonds"] = '\u2666'
	SuitMap["Clubs"] = '\u2663'
	SuitMap["Hearts"] = '\u2665'
	for i := 0; i < len(deck); i++ { // Display the cards
		if i == 0 && dealerFlag == true {
			fmt.Println("? '?'") // Hide dealer's first card
		} else {
			fmt.Printf("%v"+" "+"%q"+"\n",
				deck[i].Symbol, SuitMap[deck[i].Suit])
		}
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

func RemoveCard(slice []Card, s int) []Card {
	return append(slice[:s], slice[s+1:]...)
}

func DrawCard(deck *[]Card, size int) Card {
	i := 0 // Top card
	thisCard := (*deck)[i]
	*deck = RemoveCard(*deck, i)
	return thisCard
}
