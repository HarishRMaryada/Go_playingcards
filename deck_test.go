package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length of 12, but we got %v", len(d))
	}
}
