package poker

import "testing"

func TestNewHand(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"Valid 5 cards", "H2 SQ C2 D2 CQ", false},
		{"Valid 7 cards", "H2 SQ C2 D2 CQ HK SA", false},
		{"Too few cards", "H2 SQ C2", true},
		{"Too many cards", "H2 SQ C2 D2 CQ HK SA H3", true},
		{"Invalid card", "H2 XX C2 D2 CQ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewHand(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewHand_DuplicateCard(t *testing.T) {
	_, err := NewHand("H2 H2 SQ C2 D2")
	if err == nil {
		t.Error("Expected error for duplicate cards, got nil")
	}
}

func TestHandEvaluate_FullHouse(t *testing.T) {
	hand, err := NewHand("H2 SQ C2 D2 CQ")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != FullHouse {
		t.Errorf("Evaluate() = %v, want %v", val, FullHouse)
	}
}

func TestHandEvaluate_RoyalFlush(t *testing.T) {
	hand, err := NewHand("HT HJ HQ HK HA")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != RoyalFlush {
		t.Errorf("Evaluate() = %v, want %v", val, RoyalFlush)
	}
}

func TestHandEvaluate_StraightFlush(t *testing.T) {
	hand, err := NewHand("H5 H6 H7 H8 H9")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != StraightFlush {
		t.Errorf("Evaluate() = %v, want %v", val, StraightFlush)
	}
}

func TestHandEvaluate_FourOfAKind(t *testing.T) {
	hand, err := NewHand("H2 S2 C2 D2 HK")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != FourOfAKind {
		t.Errorf("Evaluate() = %v, want %v", val, FourOfAKind)
	}
}

func TestHandEvaluate_Flush(t *testing.T) {
	hand, err := NewHand("H2 H5 H7 H9 HK")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != Flush {
		t.Errorf("Evaluate() = %v, want %v", val, Flush)
	}
}

func TestHandEvaluate_Straight(t *testing.T) {
	hand, err := NewHand("H5 S6 C7 D8 H9")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != Straight {
		t.Errorf("Evaluate() = %v, want %v", val, Straight)
	}
}

func TestHandEvaluate_StraightWheel(t *testing.T) {
	hand, err := NewHand("HA H2 S3 C4 D5")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != Straight {
		t.Errorf("Evaluate() = %v, want %v", val, Straight)
	}
}

func TestHandEvaluate_ThreeOfAKind(t *testing.T) {
	hand, err := NewHand("H2 S2 C2 D5 HK")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != ThreeOfAKind {
		t.Errorf("Evaluate() = %v, want %v", val, ThreeOfAKind)
	}
}

func TestHandEvaluate_TwoPair(t *testing.T) {
	hand, err := NewHand("H2 S2 CK DK H5")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != TwoPair {
		t.Errorf("Evaluate() = %v, want %v", val, TwoPair)
	}
}

func TestHandEvaluate_OnePair(t *testing.T) {
	hand, err := NewHand("H2 S2 C5 D7 HK")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != OnePair {
		t.Errorf("Evaluate() = %v, want %v", val, OnePair)
	}
}

func TestHandEvaluate_HighCard(t *testing.T) {
	hand, err := NewHand("H2 S5 C7 D9 HK")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != HighCard {
		t.Errorf("Evaluate() = %v, want %v", val, HighCard)
	}
}

func TestHandEvaluate_SevenCards(t *testing.T) {
	// 7 Karten mit einem Flush darin
	hand, err := NewHand("H2 H5 H7 H9 HK S3 C4")
	if err != nil {
		t.Fatalf("NewHand() error = %v", err)
	}

	val, _ := hand.Evaluate()
	if val != Flush {
		t.Errorf("Evaluate() = %v, want %v", val, Flush)
	}
}

func TestHandCompare_Win(t *testing.T) {
	hand1, _ := NewHand("H2 SQ C2 D2 CQ") // Full House
	hand2, _ := NewHand("H5 S6 C7 D8 H9") // Straight

	result := hand1.Compare(hand2)
	if result != 1 {
		t.Errorf("Compare() = %v, want 1", result)
	}
}

func TestHandCompare_Lose(t *testing.T) {
	hand1, _ := NewHand("H5 S6 C7 D8 H9") // Straight
	hand2, _ := NewHand("H2 SQ C2 D2 CQ") // Full House

	result := hand1.Compare(hand2)
	if result != -1 {
		t.Errorf("Compare() = %v, want -1", result)
	}
}

func TestHandCompare_Tie(t *testing.T) {
	hand1, _ := NewHand("H5 S6 C7 D8 H9") // Straight to 9
	hand2, _ := NewHand("C5 D6 H7 S8 C9") // Straight to 9

	result := hand1.Compare(hand2)
	if result != 0 {
		t.Errorf("Compare() = %v, want 0", result)
	}
}

func TestHandCompare_SameTypeHigherKicker(t *testing.T) {
	hand1, _ := NewHand("HA SK C5 D7 H9") // High Card Ace
	hand2, _ := NewHand("HK SQ C5 D7 H9") // High Card King

	result := hand1.Compare(hand2)
	if result != 1 {
		t.Errorf("Compare() = %v, want 1", result)
	}
}

func TestIsFlush(t *testing.T) {
	hand1, _ := NewHand("H2 H5 H7 H9 HK")
	if !hand1.isFlush() {
		t.Error("isFlush() = false, want true")
	}

	hand2, _ := NewHand("H2 S5 H7 H9 HK")
	if hand2.isFlush() {
		t.Error("isFlush() = true, want false")
	}
}

func TestIsStraight(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		want     bool
		highCard Rank
	}{
		{"Normal straight", "H5 S6 C7 D8 H9", true, Nine},
		{"Wheel (A-2-3-4-5)", "HA H2 S3 C4 D5", true, Five},
		{"Not a straight", "H2 S5 C7 D9 HK", false, 0},
		{"High straight", "HT SJ CQ DK HA", true, Ace},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hand, _ := NewHand(tt.input)
			got, highCard := hand.isStraight()
			if got != tt.want {
				t.Errorf("isStraight() got = %v, want %v", got, tt.want)
			}
			if got && highCard != tt.highCard {
				t.Errorf("isStraight() highCard = %v, want %v", highCard, tt.highCard)
			}
		})
	}
}

func TestGetRankCounts(t *testing.T) {
	hand, _ := NewHand("H2 S2 C2 DK HK")
	counts := hand.getRankCounts()

	if counts[Two] != 3 {
		t.Errorf("getRankCounts()[Two] = %v, want 3", counts[Two])
	}
	if counts[King] != 2 {
		t.Errorf("getRankCounts()[King] = %v, want 2", counts[King])
	}
}

func TestHandValString(t *testing.T) {
	tests := []struct {
		val  HandVal
		want string
	}{
		{HighCard, "High Card"},
		{OnePair, "One Pair"},
		{TwoPair, "Two Pair"},
		{FullHouse, "Full House"},
		{RoyalFlush, "Royal Flush"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.val.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare_TwoPairHigherSecondPair(t *testing.T) {
	h1, _ := NewHand("H2 S2 CK DK H5") // 2,2,K,K,5
	h2, _ := NewHand("H2 S2 CQ DQ H5") // 2,2,Q,Q,5

	if h1.Compare(h2) != 1 {
		t.Error("Expected hand1 to win with higher second pair")
	}
}

func TestCompare_FullHouseHigherTrip(t *testing.T) {
	h1, _ := NewHand("H3 S3 C3 D2 H2") // 3 over 2
	h2, _ := NewHand("H4 S4 C4 D2 H2") // 4 over 2

	if h2.Compare(h1) != 1 {
		t.Error("Expected hand2 to win with higher trip")
	}
}

func TestCompare_StraightHigher(t *testing.T) {
	h1, _ := NewHand("H5 S6 C7 D8 H9")
	h2, _ := NewHand("H6 S7 C8 D9 HT")

	if h2.Compare(h1) != 1 {
		t.Error("Expected higher straight to win")
	}
}
