package main

func main() {
	// var card string = "ace of spades"
	// test := 20
	// card1 := "gini juga bisa"
	// fmt.Println(card)
	// fmt.Println(card1)
	// card1 = "changed"
	// fmt.Println(test)
	// fmt.Println(card1)
	// newCard1 := NewCard()
	// fmt.Println(newCard1)
	// numberCard:= NumberCard();
	// fmt.Println(numberCard)

	// //slice
	// cards := []string{"ace of diamonds", NewCard()}
	// cards = append(cards, "six of spades")
	// for i, card:= range cards{
	// 	fmt.Println(i,card)
	// }

	// cards1 := deck{"ace of diamonds", NewCard()}
	// cards1 = append(cards1, "six of spades")
	// cards1.print()

	// cards2 := deck{"ace of diamonds", NewCard()}
	// cards2 = append(cards1, "six of spades")
	// cards2.print2()

	cardsList := deck.NewDeck()
	cardList.print()

}

func NewCard() string {
	return "Five of Diamonds"
}
func NumberCard() int {
	return 4
}
