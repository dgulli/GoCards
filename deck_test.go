package main

import "testing"

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
