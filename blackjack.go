package main

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
	return 0
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

type Card struct {
	suit    Suit
	variant SuitVariant
}
