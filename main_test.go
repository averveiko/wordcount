package main

import "testing"

func TestGetWordCount(t *testing.T) {
	got := getWordCount("one two three")
	want := 3
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestGetWordCountEmpty(t *testing.T) {
	got := getWordCount("")
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
