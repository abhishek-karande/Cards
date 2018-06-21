package main

func main() {
	//var card string = "Ace of Spades"
	//card := newCard()
	//card = "King of Hearts"
	//fmt.Println(card)
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
