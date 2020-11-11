package casino

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDeckFull(t *testing.T) {
	deck := CreateDeck()
	assert.Equal(t, 52, len(deck.Deck), "Casino: standard card deck if full")
}

func TestIsFullOfCards(t *testing.T) {
	deck := CreateDeck()
	cardtype := Card{}
	for _, cardindeck := range deck.Deck {
		assert.IsType(t, cardtype, cardindeck, "Casino: every card is of type card")
	}
}

func TestDeckHasAllPropertiesInPlace(t *testing.T) {
	deck := CreateDeck()
	for _, cardindeck := range deck.Deck {
		assert.GreaterOrEqual(t, cardindeck.Color, 0, "Casion: card has too small color value")
		assert.NotEmpty(t, cardindeck.Face, "Casion: card has no face value")
		assert.NotEmpty(t, cardindeck.Displayname, "Casion: card has no displayname value")
		assert.GreaterOrEqual(t, cardindeck.Number, 0, "Casion: card has too small number value")
	}
}

func TestDeckForNotRepeatingCards(t *testing.T) {
	deck := CreateDeck()
	var testdeck [52]Card
	copy(testdeck[:], deck.Deck[:])
	for i, testcard := range deck.Deck {
		oneTouch := false
		for y, cardtomatch := range testdeck {
			if testcard == cardtomatch && !oneTouch {
				oneTouch = true
			} else if testcard == cardtomatch && oneTouch {
				log.Println(testcard, i, y)
				assert.FailNow(t, "Casion: one of the cards repeats in deck")

			}
		}
	}
}

func TestDeckForRandomness(t *testing.T) {
	deck1 := CreateDeck()
	deck2 := CreateDeck()
	assert.NotEqual(t, deck1, deck2, "Casino: cards don't shuffle on new deck")
}

func TestCreatePokerDeckGiveCard(t *testing.T) {
	poker := PokerDeck{}
	poker.CreatePokerDeck()
	card := Card{}
	assert.IsType(t, card, poker.CardsToPlay[0], "Casion: poker deck has cards")
	assert.Equal(t, 52, len(poker.CardsToPlay))
}

func TestDrawCardForPoker(t *testing.T) {
	poker := PokerDeck{}
	poker.CreatePokerDeck()
	cardtotest, _ := poker.DrawCardForPoker()
	card := Card{}
	assert.IsType(t, cardtotest, card, "Casino: card is of wrong type")
	assert.Equal(t, 51, len(poker.CardsToPlay), "Casino: card was not removed from deck still 52")
	for _, indeck := range poker.CardsToPlay {
		if cardtotest == indeck {
			assert.FailNow(t, "Card was not removed from deck, the same card exist")
		}
	}
	for i := 0; i < 52; i++ {
		if len(poker.CardsToPlay) <= 0 {
			_, err := poker.DrawCardForPoker()
			assert.Error(t, err, "Casino: deck does not throw an error when has no cards left to play")
		}
		poker.DrawCardForPoker()
	}
	if len(poker.CardsToPlay) > 0 {
		assert.FailNow(t, "Casino: not checked all cards")
	}
}
