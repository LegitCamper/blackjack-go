package main

import (
	"fmt"
	"math/rand/v2"
)

type Suit int64

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

func (s Suit) String() string {
	switch s {
	case Clubs:
		return "clubs"
	case Diamonds:
		return "diamonds"
	case Hearts:
		return "hearts"
	case Spades:
		return "spades"
	}
	return "unknown"
}

func randomSuit() Suit {
	suit := rand.Int()
	switch suit {
	case 0:
		return Clubs
	case 1:
		return Diamonds
	case 2:
		return Hearts
	case 3:
		return Spades
	default:
		return Clubs
	}
}

type SuitVariant int64

const (
	Ace SuitVariant = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func (s SuitVariant) value() int64 {
	switch s {
	case Ace:
		return 11
	case Two:
		return 2
	case Three:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	case Ten:
		return 10
	case Jack:
		return 10
	case Queen:
		return 10
	case King:
		return 10
	}
	return 11
}

func (s SuitVariant) String() string {
	switch s {
	case Ace:
		return "ace"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "jack"
	case Queen:
		return "queen"
	case King:
		return "king"
	}
	return "unknown"
}

func randomSuitVariant() SuitVariant {
	variant := rand.IntN(12)
	switch variant {
	case 0:
		return Ace
	case 1:
		return Two
	case 2:
		return Three
	case 3:
		return Four
	case 4:
		return Five
	case 5:
		return Six
	case 6:
		return Seven
	case 7:
		return Eight
	case 8:
		return Nine
	case 9:
		return Ten
	case 10:
		return Jack
	case 11:
		return Queen
	case 12:
		return King
	default:
		panic("unknown card")
	}
}

type Card struct {
	suit    Suit
	variant SuitVariant
}

func (c Card) file() string {
	return fmt.Sprintf("/assets/%s_%s.svg", c.suit.String(), c.variant.String())
}

func randomCard() Card {
	return Card{
		suit:    randomSuit(),
		variant: randomSuitVariant(),
	}
}

func deck() []Card {
	suits := []Suit{Clubs, Diamonds, Hearts, Spades}
	variants := []SuitVariant{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, King}

	var deck []Card
	for _, suit := range suits {
		for _, variant := range variants {
			deck = append(deck, Card{
				suit:    suit,
				variant: variant,
			})
		}
	}

	return deck
}

type Shoe struct {
	cards []Card
}

func (s *Shoe) remove() Card {
	if len(s.cards) > 0 {
		card := s.cards[len(s.cards)-1]
		s.cards = s.cards[:len(s.cards)-1]
		return card
	} else {
		panic("not enough cards")
	}
}

func newShoe(decks int) Shoe {
	cards := []Card{}

	for d := 0; d < decks; d++ {
		cards = append(cards, deck()...)
	}

	// shuffle
	for i := range cards {
		j := rand.IntN(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return Shoe{
		cards: cards,
	}
}

type Hand struct {
	cards []Card
}

func (h *Hand) hit(shoe *Shoe) {
	h.cards = append(h.cards, shoe.remove())
}

func (h *Hand) deal(shoe *Shoe) {
	card1 := shoe.remove()
	card2 := shoe.remove()
	h.cards = []Card{card1, card2}
}

func new_hand() Hand {
	return Hand{
		cards: []Card{},
	}
}
