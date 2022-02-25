package main

import (
	"os"
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
	filename := "_deckTesting"
	os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}
	os.Remove(filename)
}
