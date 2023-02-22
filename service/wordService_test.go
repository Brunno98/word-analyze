package service

import (
	"testing"

	"golang.org/x/exp/maps"
)

func TestAnalyzeTextWithEmptyText(t *testing.T) {
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
		t.Error("Outcome map isn't equal to expected!")
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
		t.Error("Outcome map isn't equal to expected!")
	}
}

func TestAnalyzeTextShouldIgnoreDifferenceBetweenCaseLetter(t *testing.T) {
	text := "foo bar FOO"
	expect := map[string]int{"foo": 2, "bar": 1}

	outcome, err := AnalyzeText(&text)
	if err != nil {
		t.Error("Analyze text should not generate an error")
	}

	if !maps.Equal(expect, outcome) {
		t.Error("Outcome map isn't equal to expected!")
	}

}
