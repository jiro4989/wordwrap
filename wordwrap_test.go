package wordwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdlib(t *testing.T) {
	assert.Equal(t, 1, len("a"))
	assert.Equal(t, 3, len("あ"))
	assert.Equal(t, 1, len([]rune("あ")))
}

func TestWordWrap(t *testing.T) {
	assert.Equal(t, []string{""}, WordWrap("", 0))
	assert.Equal(t, []string{""}, WordWrap("", 1))
	assert.Equal(t, []string{"a", "b"}, WordWrap("a b", 0))
	assert.Equal(t, []string{"a", "b"}, WordWrap("a b", 1))
	assert.Equal(t, []string{"a", "b"}, WordWrap("a b", 2))
	assert.Equal(t, []string{"a", "", "b"}, WordWrap("a\n\nb", 2))
	assert.Equal(t, []string{"a b"}, WordWrap("a b", 3))
	assert.Equal(t, []string{"a b", "c"}, WordWrap("a b c", 3))
	assert.Equal(t, []string{"a", "b c"}, WordWrap("a\nb c", 3))
	assert.Equal(t, []string{"a", "b", "c"}, WordWrap("a\nb\nc", 0))
	assert.Equal(t, []string{"a a", "b b", "c c"}, WordWrap("a a b b c c", 3))
	assert.Equal(t, []string{"hello", "world"}, WordWrap("hello world", 5))
	assert.Equal(t, []string{"hello", "world"}, WordWrap("hello world", 6))
	assert.Equal(t, []string{"hello world"}, WordWrap("hello world", 7))
	assert.Equal(t, []string{"hello\tworld"}, WordWrap("hello\tworld", 7))
	assert.Equal(t, []string{"こんにち", "ワールド"}, WordWrap("こんにち ワールド", 8))
	assert.Equal(t, []string{"こんにち", "ワールド"}, WordWrap("こんにち ワールド", 9))
	assert.Equal(t, []string{"こんにち ワールド"}, WordWrap("こんにち ワールド", 10))
}

func TestIsHalfRune(t *testing.T) {
	assert.Equal(t, true, IsHalfChar('a'))
	assert.Equal(t, true, IsHalfChar('Z'))
	assert.Equal(t, true, IsHalfChar('0'))
	assert.Equal(t, true, IsHalfChar(' '))
	assert.Equal(t, true, IsHalfChar('\t'))
	assert.Equal(t, true, IsHalfChar('\n'))
	assert.Equal(t, false, IsHalfChar('あ'))
	assert.Equal(t, false, IsHalfChar('１'))
	assert.Equal(t, false, IsHalfChar('！'))
	assert.Equal(t, false, IsHalfChar('　'))
}
