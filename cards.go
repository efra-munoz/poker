package cards

// The possible ranks of a card.
type Rank int

const (
	Two Rank = 2
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

// The possible suits of a card.
type Suit int

const (
	Clubs Suit = iota
	Hearts
	Spades
	Diamonds
)

// Representation of a card with rank and suit.
type Card struct {
	rank Rank
	suit Suit
}

// Create a new card
func New(rank Rank, suit Suit) Card {
	return Card{rank, suit}
}
