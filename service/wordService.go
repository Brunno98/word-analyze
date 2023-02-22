package service

import (
	"errors"
	"strings"
)

func AnalyzeText(text *string) (map[string]int, error) {
	if len(strings.Trim(*text, " ")) == 0 {
		return nil, errors.New("Empty text")
	}

	words := strings.Split(*text, " ")
	wordsFrequency := map[string]int{}
	for _, word := range words {
		word = strings.ReplaceAll(word, ".", "")
		word = strings.ToLower(word)
		if val, ok := wordsFrequency[word]; ok {
			wordsFrequency[word] = val + 1
		} else {
			wordsFrequency[word] = 1
		}
	}

	return wordsFrequency, nil
}
