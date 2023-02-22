package service

import (
	"testing"

	"golang.org/x/exp/maps"
)

func TestAnalyzeTextWithEmptyTextShouldReturnError(t *testing.T) {
	text := ""
	_, err := AnalyzeText(&text)
	if err == nil {
		t.Error("empty text should return error")
	}

	text = "          "
	_, err = AnalyzeText(&text)
	if err == nil {
		t.Error("empty text should return error")
	}
}

func TestAnalyzeTextShouldReturnAMap(t *testing.T) {
	text := "foo bar foo"
	expect := map[string]int{"foo": 2, "bar": 1}

	outcome, err := AnalyzeText(&text)
	if err != nil {
		t.Error("Analyze text should not generate an error")
	}

	if !maps.Equal(expect, outcome) {
		t.Errorf("Outcome map isn't equal to expected! expected: %v , but received: %v", expect, outcome)
	}
}

func TestAnalyzeTextShouldIgnoreDotCaracter(t *testing.T) {

	text := "foo bar foo."
	expect := map[string]int{"foo": 2, "bar": 1}

	outcome, err := AnalyzeText(&text)
	if err != nil {
		t.Error("Analyze text should not generate an error")
	}

	if !maps.Equal(expect, outcome) {
		t.Errorf("Outcome map isn't equal to expected! expected: %v but received: %v", expect, outcome)
	}
}

func TestAnalyzeTextShouldIgnoreDifferenceBetweenCaseLetter(t *testing.T) {
	text := "foo bar FOO"
	expect := map[string]int{"foo": 2, "bar": 1}

	outcome, err := AnalyzeText(&text)
	if err != nil {
		t.Errorf("Analyze text should not generate an error expected: %v | received: %v", expect, outcome)
	}

	if !maps.Equal(expect, outcome) {
		t.Errorf("Outcome map isn't equal to expected! expected: %v | received: %v", expect, outcome)
	}

}

func TestGetMostFrequentWord(t *testing.T) {
	words := map[string]int{"zzz": 1, "vvv": 2, "foo": 2, "bar": 1}
	expect := map[string]int{"foo": 2}

	outcome := GetMostFrequentWord(&words)

	if !maps.Equal(expect, *outcome) {
		t.Errorf("Outcome map isn't equal to expected! expected: %v | received: %v", expect, *outcome)
	}
}

func TestGetLessFrequentWord(t *testing.T) {
	words := map[string]int{"foo": 2, "zzz": 1, "aaa": 1, "bar": 1}
	expect := map[string]int{"aaa": 1}

	outcome := GetLessFrequentWord(&words)

	if !maps.Equal(expect, *outcome) {
		t.Errorf("Outcome map isn't equal to expected! expected: %v | received: %v", expect, *outcome)
	}
}

func TestGetMostFromEmptyMapShouldRenturnNil(t *testing.T) {
	var emptyMap map[string]int
	mostFrequency := GetMostFrequentWord(&emptyMap)

	if mostFrequency != nil {
		t.Errorf("GetMostFrequency return a value when given a empty map!")
	}

}

func TestGetLessFrequencyFromEmptyMapShouldRenturnNil(t *testing.T) {
	var emptyMap map[string]int
	lessFrequency := GetLessFrequentWord(&emptyMap)

	if lessFrequency != nil {
		t.Errorf("GetLessFrequency return a value when given a empty map!")
	}
}
