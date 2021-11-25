package languagesstore

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestingNew(t *testing.T) {
	t.Run("success", func(t *testing.T){
		assert := assert.New(t)
		languagesStore := New()
		assert.NotNil(t, languagesStore )
	})

	t.Run("failure", func(t *testing.T){
		assert := assert.New(t)
		assert.Nil(t, nil )
	})

}


func TestingLanguages(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		store := LanguagesStore{map[string]string{}}
		array := store.Languages()
		assert.Equal(t, len(array), 0, "they should be equal")
	})
}
