package model_test

import (
	"testing"

	"github.com/brunno98/word-analyze/model"
)

func TestCreateWordsFromTextMustReturnAnArrayOfString(t *testing.T) {
	text, _ := model.NewText("Foo Bar")
	expecteds := []string{"foo", "bar"}

	words, err := model.NewWords(text)
	if err != nil {
		t.Errorf("Error while creating words from text [%s]", text)
	}

	for index, word := range *words {
		expected := expecteds[index]
		if word != expected {
			t.Errorf("word diferent from expected! expected: [%s], received: [%s]", expected, word)
		}
	}
}

func TestWordsCreatedMustNotHaveCommaAtEnd(t *testing.T) {
	text, _ := model.NewText("Foo, Bar")
	expecteds := []string{"foo", "bar"}

	words, err := model.NewWords(text)
	if err != nil {
		t.Errorf("Error while creating words from text [%s]", text)
	}

	for index, word := range *words {
		expected := expecteds[index]
		if word != expected {
			t.Errorf("word diferent from expected! expected: [%s], received: [%s]", expected, word)
		}
	}
}

func TestWordsCreatedMustNotHaveDotAtEnd(t *testing.T) {
	text, _ := model.NewText("Foo. Bar")
	expecteds := []string{"foo", "bar"}

	words, err := model.NewWords(text)
	if err != nil {
		t.Errorf("Error while creating words from text [%s]", text)
	}

	for index, word := range *words {
		expected := expecteds[index]
		if word != expected {
			t.Errorf("word diferent from expected! expected: [%s], received: [%s]", expected, word)
		}
	}
}
