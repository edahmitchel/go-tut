package main

func main() {
	// cards := deck{"1,2,3,4"}
	// cards.print()
	cards := newDeckFromFIle("deck")
	cards.print()
	cards.saveToFile("deck")
}
