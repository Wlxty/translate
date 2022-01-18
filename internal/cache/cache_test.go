package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetCache(t *testing.T) {
	var rt = Through{MemoryCache: NewInMemoryCache()}
	duration := time.Hour * 2
	value, _ := rt.Get("sample", Sample, duration)
	assert.Equalf(t, value, "Sample", "They should be equal")

}

func Sample() (interface{}, error) {
	return "Sample", nil
}
