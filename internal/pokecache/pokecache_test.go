package pokecache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Hour * 100)
	assert.NotNil(t, cache.cache)
}

func TestAddCache(t *testing.T) {
	cache := NewCache(time.Hour * 100)

	tests := []struct {
		input struct {
			key   string
			value []byte
		}
		expected struct {
			value []byte
			ok    bool
		}
	}{
		{
			input: struct {
				key   string
				value []byte
			}{
				key:   "key1",
				value: []byte("value1"),
			},
			expected: struct {
				value []byte
				ok    bool
			}{
				value: []byte("value1"),
				ok:    true,
			},
		},
	}

	for _, test := range tests {
		cache.Add(test.input.key, test.input.value)
		actual, ok := cache.Get(test.input.key)
		assert.Equal(t, test.expected.value, actual)
		assert.Equal(t, test.expected.ok, ok)
	}
}

func TestReap(t *testing.T) {
	cache := NewCache(time.Millisecond * 1)
	cache.Lock()
	cache.cache["key1"] = cacheEntry{val: []byte("value1"), createdAt: time.Now().UTC().Add(-time.Millisecond * 3)}
	cache.cache["key2"] = cacheEntry{val: []byte("value2"), createdAt: time.Now().UTC().Add(-time.Millisecond * 2)}
	cache.cache["key3"] = cacheEntry{val: []byte("value3"), createdAt: time.Now().UTC().Add(-time.Millisecond * 1)}
	cache.cache["key4"] = cacheEntry{val: []byte("value4"), createdAt: time.Now().UTC().Add(time.Millisecond * 1)}
	cache.cache["key5"] = cacheEntry{val: []byte("value5"), createdAt: time.Now().UTC().Add(time.Millisecond * 2)}
	cache.cache["key6"] = cacheEntry{val: []byte("value6"), createdAt: time.Now().UTC().Add(time.Millisecond * 3)}
	cache.Unlock()

	time.Sleep(time.Second * 1)

	cache.Lock()
	assert.Equal(t, 3, len(cache.cache))
	cache.Unlock()
}
