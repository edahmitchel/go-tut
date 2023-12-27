package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		{
			fmt.Println(i, card)
		}

	}
}
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) toString() string {
	return strings.Join(d, ",")
}
func (d deck) saveTofile(file string) error {
	return os.WriteFile(file, []byte(d.toString()), 0666)

}
func newDeckFromFIle(file string) deck {
	byteSlice, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return newDeck()
		// os.Exit(1)
	}
	s := strings.Split(string(byteSlice), ",")
	d := deck(s)
	return d

}
