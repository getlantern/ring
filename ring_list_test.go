package ring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroCapacity(t *testing.T) {
	l := NewList(0)
	l.Push("a")
	assert.Equal(t, 1, l.Len())
}

func TestRingList(t *testing.T) {
	l := NewList(2)

	checkContents := func(expectedLen int, expectedForward string, expectedBackward string) {
		assert.Equal(t, expectedLen, l.Len())
		if expectedLen == 0 {
			return
		}

		// complete iteration
		append, word := wordBuilder()
		l.IterateForward(func(letter interface{}) bool {
			append(letter)
			return true
		})
		assert.Equal(t, expectedForward, word())
		append, word = wordBuilder()
		l.IterateBackward(func(letter interface{}) bool {
			append(letter)
			return true
		})
		assert.Equal(t, expectedBackward, word())

		// short iteration
		append, word = wordBuilder()
		l.IterateForward(func(letter interface{}) bool {
			append(letter)
			return false
		})
		assert.Equal(t, string(rune(expectedForward[0])), word())
		append, word = wordBuilder()
		l.IterateBackward(func(letter interface{}) bool {
			append(letter)
			return false
		})
		assert.Equal(t, string(rune(expectedBackward[0])), word())
	}

	// Empty buffer
	checkContents(0, "", "")

	// Single element
	l.Push("a")
	checkContents(1, "a", "a")

	// Full buffer
	l.Push("b")
	checkContents(2, "ab", "ba")

	// Wrapped buffer
	l.Push("c")
	checkContents(2, "bc", "cb")

	// Wraped twice buffer
	l.Push("d")
	l.Push("e")
	checkContents(2, "de", "ed")
}

func wordBuilder() (func(letter interface{}), func() string) {
	word := ""
	return func(letter interface{}) {
			word += letter.(string)
		}, func() string {
			return word
		}
}
