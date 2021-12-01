package translated

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestingNew(t *testing.T) {
	t.Run("success", func(t *testing.T){
		assert := assert.New(t)
		repository := New()
		assert.NotNil(t, repository )
	})

	t.Run("failure", func(t *testing.T){
		assert := assert.New(t)
		assert.Nil(t, nil )
	})

}