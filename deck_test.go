package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	//test for expected amounts of cards
	if len(d) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(d))

	}
	//test for expected first card in deck
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the Ace of Spades but got %v", d[0])
	}
	//test for expected last card in deck
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs, but got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
