package containers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {

	t.Run("test set - insert", func(t *testing.T) {
		s := CreateSet[string]()
		s.Insert("a")
		_, ok := s.set["a"]
		assert.True(t, ok)
	})

	t.Run("test set - contains", func(t *testing.T) {
		s := CreateSet[string]()
		s.Insert("a")
		assert.True(t, s.Contains("a"))
	})

	t.Run("test set - remvove", func(t *testing.T) {
		s := CreateSet[string]()
		s.Insert("a")
		assert.True(t, s.Contains("a"))
		s.Remove("a")
		assert.False(t, s.Contains("a"))
	})
}
