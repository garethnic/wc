package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	words := bytes.NewBufferString("words1 words2 words3")

	exp := 3

	res := count(words, false, false)

	if exp != res {
		t.Errorf("Expected %d but got %d\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	words := bytes.NewBufferString("words1 words2\n words3")

	exp := 2

	res := count(words, true, false)

	if exp != res {
		t.Errorf("Expected %d but got %d\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	words := bytes.NewBufferString("words1 words2 words3")

	exp := words.Len()

	res := count(words, false, true)

	if exp != res {
		t.Errorf("Expected %d but got %d\n", exp, res)
	}
}
