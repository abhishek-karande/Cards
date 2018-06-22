package main

func main() {

	cards := newDeck()
	//fmt.Println(cards.toString())
	cards.saveToFile("myCards")
	cards = newDeckFromFile("myCards")
	cards.Shuffle()
	cards.print()
}
