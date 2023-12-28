package main

import (
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 15 {
		t.Errorf("Expected deck length of 15, but got %d", len(d))
	}
}
func TestDeal(t *testing.T) {
	d := deck{"Card 1", "Card 2", "Card 3", "Card 4", "Card 5"}
	hand, remaining := deal(d, 2)

	expectedHand := deck{"Card 1", "Card 2"}
	expectedRemaining := deck{"Card 3", "Card 4", "Card 5"}

	if !reflect.DeepEqual(hand, expectedHand) {
		t.Errorf("Expected hand: %v, got: %v", expectedHand, hand)
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("Expected remaining: %v, got: %v", expectedRemaining, remaining)
	}
}
func TestDeckToString(t *testing.T) {
	d := deck{"Card 1", "Card 2", "Card 3"}
	expected := "Card 1,Card 2,Card 3"

	result := d.toString()

	if result != expected {
		t.Errorf("Expected: %s, got: %s", expected, result)
	}
}
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "test_deck.txt"
	defer func() {
		// Clean up after the test
		err := os.Remove(filename)
		if err != nil {
			t.Errorf("Error removing test file: %v", err)
		}
	}()

	d := deck{"Card 1", "Card 2", "Card 3"}

	// Test save to file
	err := d.saveToFile(filename)
	if err != nil {
		t.Errorf("Error saving to file: %v", err)
	}

	// Test new deck from file
	loadedDeck := newDeckFromFIle(filename)

	if !reflect.DeepEqual(d, loadedDeck) {
		t.Errorf("Expected deck: %v, got: %v", d, loadedDeck)
	}
}
