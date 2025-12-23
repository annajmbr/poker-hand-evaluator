package main

import (
	"fmt"
	"os"

	"github.com/yourusername/poker-hand-evaluator/poker"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "eval":
		if len(os.Args) < 3 {
			fmt.Println("Fehler: Hand erforderlich")
			printUsage()
			return
		}
		evaluateHand(os.Args[2])

	case "compare":
		if len(os.Args) < 4 {
			fmt.Println("Fehler: Zwei Hände erforderlich")
			printUsage()
			return
		}
		compareHands(os.Args[2], os.Args[3])

	case "best":
		if len(os.Args) < 3 {
			fmt.Println("Fehler: 7 Karten erforderlich")
			printUsage()
			return
		}
		findBestHand(os.Args[2])

	default:
		fmt.Printf("Unbekannter Befehl: %s\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Texas Hold'em Poker Hand Evaluator")
	fmt.Println("\nVerwendung:")
	fmt.Println("  poker eval <hand>         - Bewertet eine 5-Karten-Hand")
	fmt.Println("  poker compare <h1> <h2>   - Vergleicht zwei Hände")
	fmt.Println("  poker best <7cards>       - Findet beste 5-Karten-Hand aus 7 Karten")
	fmt.Println("\nKartenformat:")
	fmt.Println("  Farbe: H(Hearts), D(Diamonds), C(Clubs), S(Spades)")
	fmt.Println("  Wert:  2-9, T(10), J(Jack), Q(Queen), K(King), A(Ace)")
	fmt.Println("\nBeispiele:")
	fmt.Println("  poker eval \"H2 SQ C2 D2 CQ\"")
	fmt.Println("  poker compare \"H2 SQ C2 D2 CQ\" \"H5 S6 C7 D8 H9\"")
	fmt.Println("  poker best \"H2 SQ C2 D2 CQ HK SA\"")
}

func evaluateHand(handStr string) {
	hand, err := poker.NewHand(handStr)
	if err != nil {
		fmt.Printf("Fehler beim Parsen der Hand: %v\n", err)
		return
	}

	val, kickers := hand.Evaluate()
	fmt.Printf("Hand: %s\n", handStr)
	fmt.Printf("Bewertung: %s\n", val)
	fmt.Printf("Kickers: ")
	for i, k := range kickers {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(k)
	}
	fmt.Println()
}

func compareHands(hand1Str, hand2Str string) {
	hand1, err := poker.NewHand(hand1Str)
	if err != nil {
		fmt.Printf("Fehler beim Parsen von Hand 1: %v\n", err)
		return
	}

	hand2, err := poker.NewHand(hand2Str)
	if err != nil {
		fmt.Printf("Fehler beim Parsen von Hand 2: %v\n", err)
		return
	}

	val1, _ := hand1.Evaluate()
	val2, _ := hand2.Evaluate()

	fmt.Printf("Hand 1: %s (%s)\n", hand1Str, val1)
	fmt.Printf("Hand 2: %s (%s)\n", hand2Str, val2)
	fmt.Println()

	result := hand1.Compare(hand2)
	switch result {
	case 1:
		fmt.Println("Ergebnis: Hand 1 gewinnt!")
	case -1:
		fmt.Println("Ergebnis: Hand 2 gewinnt!")
	case 0:
		fmt.Println("Ergebnis: Gleichstand!")
	}
}

func findBestHand(sevenCardsStr string) {
	hand, err := poker.NewHand(sevenCardsStr)
	if err != nil {
		fmt.Printf("Fehler beim Parsen der Hand: %v\n", err)
		return
	}

	if len(hand.Cards) != 7 {
		fmt.Printf("Fehler: Genau 7 Karten erforderlich, erhalten: %d\n", len(hand.Cards))
		return
	}

	val, kickers := hand.Evaluate()
	fmt.Printf("7 Karten: %s\n", sevenCardsStr)
	fmt.Printf("Beste Hand: %s\n", val)
	fmt.Printf("Kickers: ")
	for i, k := range kickers {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(k)
	}
	fmt.Println()
}
