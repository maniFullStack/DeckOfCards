package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length 16, but got %v", len(d))

	}

	if d[0] != "Ace of Spades" {

		t.Errorf("Expected, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clover" {

		t.Errorf("Expected, but got %v", d[len(d)-1])
	}
}
