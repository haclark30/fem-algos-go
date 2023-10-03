package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLru(t *testing.T) {
	lru := NewLruCache[string, int](3)
	assert.Nil(t, lru.Get("foo"))
	lru.Update("foo", 69)
	assert.Equal(t, 69, *lru.Get("foo"))
	lru.Update("bar", 420)
	assert.Equal(t, 420, *lru.Get("bar"))
	lru.Update("baz", 1337)
	assert.Equal(t, 1337, *lru.Get("baz"))

	lru.Update("ball", 69420)
	assert.Equal(t, 69420, *lru.Get("ball"))
	assert.Nil(t, lru.Get("foo"))
	assert.Equal(t, 420, *lru.Get("bar"))
	lru.Update("foo", 69)
	assert.Equal(t, 420, *lru.Get("bar"))
	assert.Equal(t, 69, *lru.Get("foo"))

	assert.Nil(t, lru.Get("baz"))

}
