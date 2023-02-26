package model

import (
	"errors"
	"strings"
)

type Text string

func NewText(text string) (*Text, error) {
	text = strings.Trim(text, " ")
	if text == "" {
		return nil, errors.New("Empty text")
	}

	text = strings.ToLower(text)

	t := Text(text)
	return &t, nil
}

func (t *Text) String() string {
	return string(*t)
}
