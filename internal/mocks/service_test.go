package mocks

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"translateapp/internal/cache"
	"translateapp/internal/libretranslate"
	"translateapp/internal/logger"
	"translateapp/internal/translateapp"
)

func TestService(t *testing.T) {
	lt := LibreTranslator{}
	languages := []translateapp.Language{
		{
			"pl",
			"polish",
		},
	}
	lt.On("GetLanguages", mock.Anything).Return(&languages, nil)
	logger := logger.NewLogger("debug", true)
	client := libretranslate.NewClient(logger, "http://libretranslate:5000/")
	rt := cache.Through{Proxy: cache.NewInMemoryProxy()}
	cached := translateapp.Cache{client, rt}
	var cacher translateapp.Cacher = &cached
	service := &translateapp.Service{
		Logger: logger,
		Cached: cacher,
	}
	value, err, cachedKey := service.Languages()
	data, err := json.Marshal(languages)
	require.NoError(t, err)
	require.Equal(t, cachedKey, "languages")
	require.Equal(t, data, value)
	lt.AssertCalled(t, "GetLanguages")
}
