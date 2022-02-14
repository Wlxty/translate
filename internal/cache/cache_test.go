package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"translateapp/internal/logger"
)

func TestGetCache(t *testing.T) {
	logger := logger.NewLogger("debug", true)
	var rt = Through{MemoryCache: NewInMemoryCache(logger)}
	duration := time.Hour * 2
	value, _ := rt.Get("sample", Sample, duration)
	assert.Equalf(t, value, "Sample", "They should be equal")

}

func Sample() (interface{}, error) {
	return "Sample", nil
}
