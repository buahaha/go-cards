package casino

import (
	"errors"
	"math/rand"
	"time"
)

const maxCardsInStandardDeck = 52

type Card struct {
	Face        string
	Number      int
	Color       int
	Displayname string
}

type DeckOfCards struct {
	Deck [52]Card
}

func CreateDeck() DeckOfCards {
	var deck []Card
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for {
		var number = r1.Intn(13)
		var color = r1.Intn(4)
		var face string
		var displayname string
		switch number {
		case 0:
			switch color {
			case 3:
				face = "🂢"
				displayname = "2 of spades"
			case 2:
				face = "🂲"
				displayname = "2 of of hearts"
			case 1:
				face = "🃒"
				displayname = "2 of clubs"
			case 0:
				face = "🃂"
				displayname = "2 of diamonds"
			}
		case 1:
			switch color {
			case 3:
				face = "🂣"
				displayname = "3 of spades"
			case 2:
				face = "🂳"
				displayname = "3 of hearts"
			case 1:
				face = "🃓"
				displayname = "3 of clubs"
			case 0:
				face = "🃃"
				displayname = "3 of diamonds"
			}
		case 2:
			switch color {
			case 3:
				face = "🂤"
				displayname = "4 of spades"
			case 2:
				face = "🂴"
				displayname = "4 of hearts"
			case 1:
				face = "🃔"
				displayname = "4 of clubs"
			case 0:
				face = "🃄"
				displayname = "4 of diamonds"
			}
		case 3:
			switch color {
			case 3:
				face = "🂥"
				displayname = "5 of spades"
			case 2:
				face = "🂵"
				displayname = "5 of hearts"
			case 1:
				face = "🃕"
				displayname = "5 of clubs"
			case 0:
				face = "🃅"
				displayname = "5 of diamonds"
			}
		case 4:
			switch color {
			case 3:
				face = "🂦"
				displayname = "6 of spades"
			case 2:
				face = "🂶"
				displayname = "6 of hearts"
			case 1:
				face = "🃖"
				displayname = "6 of clubs"
			case 0:
				face = "🃆"
				displayname = "6 of diamonds"
			}
		case 5:
			switch color {
			case 3:
				face = "🂧"
				displayname = "7 of spades"
			case 2:
				face = "🂷"
				displayname = "7 of hearts"
			case 1:
				face = "🃗"
				displayname = "7 of clubs"
			case 0:
				face = "🃇"
				displayname = "7 of diamonds"
			}
		case 6:
			switch color {
			case 3:
				face = "🂨"
				displayname = "8 of spades"
			case 2:
				face = "🂸"
				displayname = "8 of hearts"
			case 1:
				face = "🃘"
				displayname = "8 of clubs"
			case 0:
				face = "🃈"
				displayname = "8 of diamonds"
			}
		case 7:
			switch color {
			case 3:
				face = "🂩"
				displayname = "9 of spades"
			case 2:
				face = "🂹"
				displayname = "9 of hearts"
			case 1:
				face = "🃙"
				displayname = "9 of clubs"
			case 0:
				face = "🃉"
				displayname = "9 of diamonds"
			}
		case 8:
			switch color {
			case 3:
				face = "🂪"
				displayname = "10 of spades"
			case 2:
				face = "🂺"
				displayname = "10 of hearts"
			case 1:
				face = "🃚"
				displayname = "10 of clubs"
			case 0:
				face = "🃊"
				displayname = "10 of diamonds"
			}
		case 9:
			switch color {
			case 3:
				face = "🂫"
				displayname = "Jack of spades"
			case 2:
				face = "🂻"
				displayname = "Jack of hearts"
			case 1:
				face = "🃛"
				displayname = "Jack of clubs"
			case 0:
				face = "🃋"
				displayname = "Jack of diamonds"
			}
		case 10:
			switch color {
			case 3:
				face = "🂭"
				displayname = "Queen of spades"
			case 2:
				face = "🂽"
				displayname = "Queen of hearts"
			case 1:
				face = "🃝"
				displayname = "Queen of clubs"
			case 0:
				face = "🃍"
				displayname = "Queen of diamonds"
			}
		case 11:
			switch color {
			case 3:
				face = "🂮"
				displayname = "King of spades"
			case 2:
				face = "🂾"
				displayname = "King of hearts"
			case 1:
				face = "🃞"
				displayname = "King of clubs"
			case 0:
				face = "🃎"
				displayname = "King of diamonds"
			}
		case 12:
			switch color {
			case 3:
				face = "🂡"
				displayname = "Ace of spades"
			case 2:
				face = "🂱"
				displayname = "Ace of hearts"
			case 1:
				face = "🃑"
				displayname = "Ace of clubs"
			case 0:
				face = "🃁"
				displayname = "Ace of diamonds"
			}
		}
		mycard := Card{
			Face:        face,
			Number:      number,
			Color:       color,
			Displayname: displayname,
		}

		if len(deck) > 0 && len(deck) < 52 {
			oneTouch := false
			for _, testcard := range deck {
				if testcard == mycard {
					oneTouch = true
					break
				}
			}
			if !oneTouch {
				deck = append(deck[:], mycard)
			}
		} else if len(deck) == 0 {
			deck = append(deck[:], mycard)
		} else {
			break
		}
	}
	mydeck := DeckOfCards{}
	copy(mydeck.Deck[:], deck[:])
	return mydeck
}

type PokerDeck struct {
	CardsToPlay []Card
}

func (p *PokerDeck) CreatePokerDeck() []Card {
	deck := CreateDeck()
	p.CardsToPlay = deck.Deck[:]
	return p.CardsToPlay
}

func (p *PokerDeck) DrawCardForPoker() (Card, error) {
	if len(p.CardsToPlay) <= 0 {
		return Card{}, errors.New("Casino: Not enough cards in deck")
	}
	card := p.CardsToPlay[0]
	p.CardsToPlay = p.CardsToPlay[1:]
	return card, nil
}
