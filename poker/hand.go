package poker

import (
	"fmt"
	"sort"
	"strings"
)

// HandVal repräsentiert die Bewertung einer Poker-Hand
type HandVal int

const (
	HighCard HandVal = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

// String gibt die String-Repräsentation der Hand-Bewertung zurück
func (hv HandVal) String() string {
	return [...]string{
		"High Card",
		"One Pair",
		"Two Pair",
		"Three of a Kind",
		"Straight",
		"Flush",
		"Full House",
		"Four of a Kind",
		"Straight Flush",
		"Royal Flush",
	}[hv]
}

// Hand repräsentiert eine Poker-Hand mit bis zu 7 Karten
type Hand struct {
	Cards []*Card
}

// NewHand erstellt eine neue Hand aus einem String (z.B. "H2 SQ C2 D2 CQ")
func NewHand(s string) (*Hand, error) {
	parts := strings.Fields(s)
	if len(parts) < 5 || len(parts) > 7 {
		return nil, fmt.Errorf("Hand muss 5-7 Karten enthalten, erhalten: %d", len(parts))
	}

	hand := &Hand{Cards: make([]*Card, 0, len(parts))}
	seen := make(map[string]bool)

	for _, part := range parts {
		card, err := NewCard(part)
		if err != nil {
			return nil, err
		}

		key := card.String()
		if seen[key] {
			return nil, fmt.Errorf("doppelte Karte in Hand: %s", key)
		}
		seen[key] = true

		hand.Cards = append(hand.Cards, card)
	}

	return hand, nil
}

// Evaluate bewertet eine Hand und gibt den HandVal zurück
func (h *Hand) Evaluate() (HandVal, []Rank) {
	if len(h.Cards) == 5 {
		return h.evaluateFive()
	}
	if len(h.Cards) == 7 {
		return h.findBestFive()
	}
	return HighCard, nil
}

// evaluateFive bewertet eine 5-Karten-Hand
func (h *Hand) evaluateFive() (HandVal, []Rank) {
	ranks := make([]Rank, len(h.Cards))
	for i, card := range h.Cards {
		ranks[i] = card.Rank
	}
	sort.Slice(ranks, func(i, j int) bool { return ranks[i] > ranks[j] })

	isFlush := h.isFlush()
	isStraight, straightHighCard := h.isStraight()

	if isFlush && isStraight {
		if straightHighCard == Ace {
			return RoyalFlush, []Rank{Ace}
		}
		return StraightFlush, []Rank{straightHighCard}
	}

	rankCounts := h.getRankCounts()
	counts := make([]int, 0)
	for _, count := range rankCounts {
		counts = append(counts, count)
	}
	sort.Slice(counts, func(i, j int) bool { return counts[i] > counts[j] })

	if counts[0] == 4 {
		return FourOfAKind, h.getKickers(rankCounts, 4)
	}

	if counts[0] == 3 && counts[1] == 2 {
		return FullHouse, h.getKickers(rankCounts, 3, 2)
	}

	if isFlush {
		return Flush, ranks
	}

	if isStraight {
		return Straight, []Rank{straightHighCard}
	}

	if counts[0] == 3 {
		return ThreeOfAKind, h.getKickers(rankCounts, 3)
	}

	if counts[0] == 2 && counts[1] == 2 {
		return TwoPair, h.getKickers(rankCounts, 2, 2)
	}

	if counts[0] == 2 {
		return OnePair, h.getKickers(rankCounts, 2)
	}

	return HighCard, ranks
}

// findBestFive findet die beste 5-Karten-Hand aus 7 Karten
func (h *Hand) findBestFive() (HandVal, []Rank) {
	if len(h.Cards) != 7 {
		return HighCard, nil
	}

	bestVal := HighCard
	var bestKickers []Rank

	// Alle Kombinationen von 5 aus 7 Karten durchprobieren
	indices := make([]int, 5)
	for i := range indices {
		indices[i] = i
	}

	for {
		subHand := &Hand{Cards: make([]*Card, 5)}
		for i, idx := range indices {
			subHand.Cards[i] = h.Cards[idx]
		}

		val, kickers := subHand.evaluateFive()
		if val > bestVal || (val == bestVal && compareKickers(kickers, bestKickers) > 0) {
			bestVal = val
			bestKickers = kickers
		}

		if !nextCombination(indices, 7) {
			break
		}
	}

	return bestVal, bestKickers
}

// isFlush prüft, ob alle Karten die gleiche Farbe haben
func (h *Hand) isFlush() bool {
	if len(h.Cards) != 5 {
		return false
	}
	suit := h.Cards[0].Suit
	for _, card := range h.Cards[1:] {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

// isStraight prüft, ob die Karten eine Straße bilden
func (h *Hand) isStraight() (bool, Rank) {
	if len(h.Cards) != 5 {
		return false, 0
	}

	ranks := make([]Rank, len(h.Cards))
	for i, card := range h.Cards {
		ranks[i] = card.Rank
	}
	sort.Slice(ranks, func(i, j int) bool { return ranks[i] < ranks[j] })

	// Normale Straße prüfen
	for i := 1; i < len(ranks); i++ {
		if ranks[i] != ranks[i-1]+1 {
			// Spezialfall: A-2-3-4-5 (Wheel)
			if i == 4 && ranks[0] == Two && ranks[1] == Three &&
				ranks[2] == Four && ranks[3] == Five && ranks[4] == Ace {
				return true, Five
			}
			return false, 0
		}
	}

	return true, ranks[4]
}

// getRankCounts zählt die Häufigkeit jedes Rangs
func (h *Hand) getRankCounts() map[Rank]int {
	counts := make(map[Rank]int)
	for _, card := range h.Cards {
		counts[card.Rank]++
	}
	return counts
}

// getKickers gibt die Kicker in absteigender Reihenfolge zurück
func (h *Hand) getKickers(rankCounts map[Rank]int, targetCounts ...int) []Rank {
	result := make([]Rank, 0)

	for _, target := range targetCounts {
		for rank := Ace; rank >= Two; rank-- {
			if rankCounts[rank] == target {
				result = append(result, rank)
			}
		}
	}

	// Rest als einzelne Kicker hinzufügen
	for rank := Ace; rank >= Two; rank-- {
		if rankCounts[rank] > 0 && !contains(result, rank) {
			for i := 0; i < rankCounts[rank]; i++ {
				result = append(result, rank)
			}
		}
	}

	return result
}

// Compare vergleicht zwei Hände und gibt zurück:
// 1 wenn h gewinnt, -1 wenn other gewinnt, 0 bei Gleichstand
func (h *Hand) Compare(other *Hand) int {
	val1, kickers1 := h.Evaluate()
	val2, kickers2 := other.Evaluate()

	if val1 > val2 {
		return 1
	}
	if val1 < val2 {
		return -1
	}

	return compareKickers(kickers1, kickers2)
}

// compareKickers vergleicht Kicker-Listen
func compareKickers(k1, k2 []Rank) int {
	minLen := len(k1)
	if len(k2) < minLen {
		minLen = len(k2)
	}

	for i := 0; i < minLen; i++ {
		if k1[i] > k2[i] {
			return 1
		}
		if k1[i] < k2[i] {
			return -1
		}
	}

	if len(k1) > len(k2) {
		return 1
	}
	if len(k1) < len(k2) {
		return -1
	}

	return 0
}

// nextCombination generiert die nächste Kombination
func nextCombination(indices []int, n int) bool {
	k := len(indices)
	for i := k - 1; i >= 0; i-- {
		if indices[i] < n-k+i {
			indices[i]++
			for j := i + 1; j < k; j++ {
				indices[j] = indices[j-1] + 1
			}
			return true
		}
	}
	return false
}

// contains prüft, ob ein Rank in einem Slice enthalten ist
func contains(slice []Rank, rank Rank) bool {
	for _, r := range slice {
		if r == rank {
			return true
		}
	}
	return false
}
