package main

import (
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

func random_suit() Suit {
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

func random_suit_variant() SuitVariant {
	variant := rand.IntN(13)
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

func random_card() Card {
	return Card{
		suit:    random_suit(),
		variant: random_suit_variant(),
	}
}
