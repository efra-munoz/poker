package cards

import "fmt"

// The possible ranks of a card.
type Rank int

const (
	Two Rank = iota + 2
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
	Ace
)

// Returns a string representation of a rank.
func (r Rank) String() string {
	return [...]string{
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eigth",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
		"Ace",
	}[r-2] // Rank enum starts at iota + 2.
}

// The possible suits of a card.
type Suit int

const (
	Clubs Suit = iota
	Hearts
	Spades
	Diamonds
)

// Returns a string representation of a suit.
func (s Suit) String() string {
	return [...]string{
		"Clubs",
		"Hearts",
		"Spades",
		"Diamonds",
	}[s]
}

// Representation of a card with rank and suit.
type Card struct {
	rank Rank
	suit Suit
}

// Create a new card
func New(rank Rank, suit Suit) Card {
	return Card{rank, suit}
}

// Returns a string replesentation af a card.
func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.rank, c.suit)
}
