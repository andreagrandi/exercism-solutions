package blackjack

var cardValues = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value, ok := cardValues[card]
	switch {
	case ok:
		return value
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
	/* If you have a pair of aces you must always split them */
	case card1 == "ace" && card2 == "ace":
		return "P"
	/* If you have a Blackjack (two cards that sum up to a value of 21),
	and the dealer does not have an ace, a figure or a ten then you automatically win.
	If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card. */
	case (ParseCard(card1)+ParseCard(card2) >= 21):
		if dealerCard != "ace" && ParseCard(dealerCard) != 10 {
			return "W"
		} else {
			return "S"
		}
	/* If your cards sum up to a value within the range [17, 20] you should always stand. */
	case ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20:
		return "S"
	/* If your cards sum up to a value within the range [12, 16] you should always stand
	unless the dealer has a 7 or higher, in which case you should always hit. */
	case ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16:
		if ParseCard(dealerCard) >= 7 {
			return "H"
		}
		return "S"
	/* If your cards sum up to 11 or lower you should always hit. */
	case ParseCard(card1)+ParseCard(card2) <= 11:
		return "H"
	default:
		return ""
	}
}
