package ringbuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroCapacity(t *testing.T) {
	rb := New(0)
	rb.Push("a")
	assert.Equal(t, 1, rb.Len())
}

func TestRingBuffer(t *testing.T) {
	rb := New(2)

	checkContents := func(expectedLen int, expectedForward string, expectedBackward string) {
		assert.Equal(t, expectedLen, rb.Len())
		if expectedLen == 0 {
			return
		}

		// complete iteration
		append, word := wordBuilder()
		rb.IterateForward(func(letter interface{}) bool {
			append(letter)
			return true
		})
		assert.Equal(t, expectedForward, word())
		append, word = wordBuilder()
		rb.IterateBackward(func(letter interface{}) bool {
			append(letter)
			return true
		})
		assert.Equal(t, expectedBackward, word())

		// short iteration
		append, word = wordBuilder()
		rb.IterateForward(func(letter interface{}) bool {
			append(letter)
			return false
		})
		assert.Equal(t, string(rune(expectedForward[0])), word())
		append, word = wordBuilder()
		rb.IterateBackward(func(letter interface{}) bool {
			append(letter)
			return false
		})
		assert.Equal(t, string(rune(expectedBackward[0])), word())
	}

	// Empty buffer
	checkContents(0, "", "")

	// Single element
	rb.Push("a")
	checkContents(1, "a", "a")

	// Full buffer
	rb.Push("b")
	checkContents(2, "ab", "ba")

	// Wrapped buffer
	rb.Push("c")
	checkContents(2, "bc", "cb")
}

func wordBuilder() (func(letter interface{}), func() string) {
	word := ""
	return func(letter interface{}) {
			word += letter.(string)
		}, func() string {
			return word
		}
}
