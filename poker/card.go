package poker

import "fmt"

// Suit repräsentiert die Farbe einer Spielkarte
type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

// String gibt die String-Repräsentation der Farbe zurück
func (s Suit) String() string {
	return [...]string{"H", "D", "C", "S"}[s]
}

// Rank repräsentiert den Wert einer Spielkarte
type Rank int

const (
	Two Rank = iota
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

// String gibt die String-Repräsentation des Werts zurück
func (r Rank) String() string {
	return [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}[r]
}

// Card repräsentiert eine Spielkarte mit Farbe und Wert
type Card struct {
	Suit Suit
	Rank Rank
}

// NewCard erstellt eine neue Karte aus einem String (z.B. "H2", "SQ")
func NewCard(s string) (*Card, error) {
	if len(s) != 2 {
		return nil, fmt.Errorf("ungültiges Kartenformat: %s", s)
	}

	var suit Suit
	switch s[0] {
	case 'H':
		suit = Hearts
	case 'D':
		suit = Diamonds
	case 'C':
		suit = Clubs
	case 'S':
		suit = Spades
	default:
		return nil, fmt.Errorf("ungültige Farbe: %c", s[0])
	}

	var rank Rank
	switch s[1] {
	case '2':
		rank = Two
	case '3':
		rank = Three
	case '4':
		rank = Four
	case '5':
		rank = Five
	case '6':
		rank = Six
	case '7':
		rank = Seven
	case '8':
		rank = Eight
	case '9':
		rank = Nine
	case 'T':
		rank = Ten
	case 'J':
		rank = Jack
	case 'Q':
		rank = Queen
	case 'K':
		rank = King
	case 'A':
		rank = Ace
	default:
		return nil, fmt.Errorf("ungültiger Rang: %c", s[1])
	}

	return &Card{Suit: suit, Rank: rank}, nil
}

// String gibt die String-Repräsentation der Karte zurück
func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.Suit, c.Rank)
}
