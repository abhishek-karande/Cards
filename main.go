package main

func main() {
	//var card string = "Ace of Spades"
	//card := newCard()
	//card = "King of Hearts"
	//fmt.Println(card)
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Giraffe of Africa")
	cards.print()
}
func newCard() string {
	return "Five of Clubs"
}
