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

func GetMostFrequentWord(words *map[string]int) *map[string]int {
	if words == nil || len(*words) == 0 {
		return nil
	}

	var mostFrequentWord string
	for k, v := range *words {
		if mostFrequentWord == "" {
			mostFrequentWord = k
			continue
		}
		if v > (*words)[mostFrequentWord] {
			mostFrequentWord = k
		} else if v == (*words)[mostFrequentWord] {
			if k < mostFrequentWord {
				mostFrequentWord = k
			}
		}
	}
	return &map[string]int{mostFrequentWord: (*words)[mostFrequentWord]}
}

func GetLessFrequentWord(words *map[string]int) *map[string]int {
	if words == nil || len(*words) == 0 {
		return nil
	}

	var lessFrequentWord string
	for k, v := range *words {
		if lessFrequentWord == "" {
			lessFrequentWord = k
			continue
		}
		if v < (*words)[lessFrequentWord] {
			lessFrequentWord = k
		} else if v == (*words)[lessFrequentWord] {
			if k < lessFrequentWord {
				lessFrequentWord = k
			}
		}
	}
	return &map[string]int{lessFrequentWord: (*words)[lessFrequentWord]}
}

func IsPalindrome(word *string) bool {
	for i := 0; i < len(*word)/2; i++ {
		if (*word)[i] != (*word)[len(*word)-1-i] {
			return false
		}
	}
	return true
}

func GetPalindromeWords(words *map[string]int) []string {
	var palindromes []string
	for word := range *words {
		if IsPalindrome(&word) {
			palindromes = append(palindromes, word)
		}
	}
	return palindromes
}
