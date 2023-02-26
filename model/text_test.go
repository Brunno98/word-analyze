package model_test

import (
	"strings"
	"testing"

	"github.com/brunno98/word-analyze/model"
)

func TestBlankSpacesBustBeRemovedAtStartAndEndOfText(t *testing.T) {
	raw_text := "       The Lazy fox       "
	text, err := model.NewText(raw_text)
	if err != nil {
		t.Errorf("Error creating a Text from value: %s. error: %s", raw_text, err)
	}

	if strings.HasPrefix(text.String(), " ") {
		t.Errorf("text %s shuldn't start with blank space", text)
	}

	if strings.HasSuffix(text.String(), " ") {
		t.Errorf("text %s shuldn't end with blank space", text)
	}

}

func TestTextShouldBeLowerCase(t *testing.T) {
	raw_text := "The Lazy FOX"

	text, err := model.NewText(raw_text)
	if err != nil {
		t.Errorf("Error creating a Text from value: %s. error: %s", raw_text, err)
	}

	if text.String() != "the lazy fox" {
		t.Errorf("Text [%s] have some upper case letters", text)
	}
}
