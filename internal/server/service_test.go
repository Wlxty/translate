package server

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"translateapp/internal/mocks"
	"translateapp/internal/translateapp"
)

func TestGetLanguages(t *testing.T) {
	lt := mocks.LibreTranslator{}
	languages := []translateapp.Language{
		{
			"pl",
			"polish",
		},
	}
	lt.On("GetLanguages", mock.Anything).Return(&languages, nil)
	_, err := json.Marshal(languages)
	require.NoError(t, err)
	val, err := lt.GetLanguages(context.Background())
	require.Equal(t, val, &languages)
}
