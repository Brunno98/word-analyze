package model

import (
	"strings"
)

type Words []string

func NewWords(text *Text) (*Words, error) {
	words := strings.Split(text.String(), " ")

	for index, word := range words {
		words[index] = removeSuffix(word, ",")
		words[index] = removeSuffix(word, ".")
	}

	w := Words(words)
	return &w, nil
}

func removeSuffix(word string, suffix string) string {
	if strings.HasSuffix(word, suffix) {
		suffixIndex := strings.LastIndex(word, suffix)
		return word[:suffixIndex]
	}
	return word
}
