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
