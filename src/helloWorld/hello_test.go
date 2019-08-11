package main

import (
	"testing"
)

func TestTalk(t *testing.T) {
	word := "hello"
	resultWord := "hello"

	result := replaceWord(word)

	if result != resultWord {
		t.Fatal("failed test")
	}
}

func TestTalk_filteringWord(t *testing.T) {
	word := "fuck you"
	resultWord := "**** you"

	result := replaceWord(word)

	if result != resultWord {
		t.Fatal("failed test")
	}
}
