package translateapp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLanguage_Languages(t *testing.T) {
	var language Language
	languages := language.Languages()
	assert.Equal(t, []Language{
		Language{"en", "English"},
		Language{"pl", "Polish"},
	}, languages)
}

func TestNewWord(t *testing.T) {
	word := NewWord("Hello")
	assert.Equal(t, Word{"Hello"}, word)
}
