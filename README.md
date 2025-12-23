# Texas Hold'em Poker Hand Evaluator

Ein vollständiges Go-Programm zur Bewertung und zum Vergleich von Poker-Händen nach den Regeln von Texas Hold'em.

---

## Funktionen

- **Hand-Bewertung**: Bewertet 5-Karten-Pokerhände
- **Hand-Vergleich**: Vergleicht zwei Pokerhände inklusive Kicker
- **Beste 5 aus 7**: Ermittelt automatisch die beste 5-Karten-Hand aus 7 Karten
- **Umfangreiche Tests**: Vollständige Unit-Tests mit hoher Code-Coverage (>90 %)

---

## Projektstruktur

```
Poker-hand-evaluator/
├── README.md
├── go.mod
├── coverage.out
├── Prompts.txt
├── PreviousPrompts.txt
├── Coverage_Poker.png
│
├── poker/
│   ├── card.go
│   ├── hand.go
│   └── hand_test.go
│
└── cmd/
    └── poker/
        └── main.go
```

---

## Installation

```bash
git clone <REPOSITORY-URL>
cd Poker-hand-evaluator
go mod tidy
```

---

## Build

```bash
go build -o poker ./cmd/poker
```

---

## Verwendung

### Hand bewerten

```bash
./poker eval "H2 SQ C2 D2 CQ"
```

**Beispielausgabe:**
```
Bewertung: Full House
```

---

### Zwei Hände vergleichen

```bash
./poker compare "H2 SQ C2 D2 CQ" "H5 S6 C7 D8 H9"
```

**Beispielausgabe:**
```
Ergebnis: Hand 1 gewinnt!
```

---

### Beste 5-Karten-Hand aus 7 Karten

```bash
./poker best "H2 H5 H7 H9 HK S3 C4"
```

**Beispielausgabe:**
```
Beste Hand: Flush
```

---

## Kartenformat

### Farben
- `H` – Hearts (Herz)
- `D` – Diamonds (Karo)
- `C` – Clubs (Kreuz)
- `S` – Spades (Pik)

### Werte
- `2–9`, `T` (10), `J`, `Q`, `K`, `A`

**Beispiele:**
- `H2` – Herz 2
- `SQ` – Pik Dame
- `CA` – Kreuz Ass

---

## Hand-Rankings (aufsteigend)

1. High Card
2. One Pair
3. Two Pair
4. Three of a Kind
5. Straight
6. Flush
7. Full House
8. Four of a Kind
9. Straight Flush
10. Royal Flush

---

## Tests

### Alle Tests ausführen

```bash
go test ./... -v
```

### Tests mit Coverage

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## Architektur

### Objektorientierter Ansatz in Go

#### Card
```go
type Card struct {
    Suit Suit
    Rank Rank
}
```

#### Hand
```go
type Hand struct {
    Cards []*Card
}
```

Alle Methoden verwenden Pointer Receiver.

---

## Lizenz

Dieses Projekt wurde im Rahmen einer universitären Lehrveranstaltung erstellt.
