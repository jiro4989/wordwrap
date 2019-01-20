package wordwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordWrap(t *testing.T) {
	assert.Equal(t, []string{""}, WordWrap("", 0))
}
