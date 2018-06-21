package main

func main() {
	//var card string = "Ace of Spades"
	//card := newCard()
	//card = "King of Hearts"
	//fmt.Println(card)
	cards := newDeck()
	cards.print()
}
func newCard() string {
	return "Five of Clubs"
}
