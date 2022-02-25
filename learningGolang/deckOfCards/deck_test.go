package main

import (
	"testing"
)

func Test_deal(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

func Test_deck_print(t *testing.T) {
}

func Test_deck_saveToFile(t *testing.T) {
}

func Test_deck_shuffle(t *testing.T) {
}

func Test_deck_toString(t *testing.T) {
}

func Test_newDeck(t *testing.T) {
}

func Test_newDeckFromFile(t *testing.T) {
}
