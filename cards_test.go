package cards

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("should create a card with given <suit> and <rank>",
		func(t *testing.T) {
			var rank Rank = Two
			var suit Suit = Spades

			card := New(rank, suit)
			expected := Card{rank, suit}

			if card != expected {
				t.Errorf("expected %q but got %q", expected, card)
			}
		})
}

func TestString(t *testing.T) {
	t.Run("should return a string representation of a Rank",
		func(t *testing.T) {
			expected := "Ace"

			if Ace.String() != expected {
				t.Errorf("expected %q but got %q", expected, Ace)
			}
		})

	t.Run("should return a string representation of a suit",
		func(t *testing.T) {
			expected := "Spades"

			if Spades.String() != expected {
				t.Errorf("expected %q but got %q", expected, Spades)
			}
		})

	t.Run("should return a string representation of a card",
		func(t *testing.T) {
			card := New(Ace, Spades)
			expected := "Ace of Spades"

			if card.String() != expected {
				t.Errorf("expected %q but got %q", expected, card)
			}
		})
}
