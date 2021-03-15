package printer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPrinter(t *testing.T) {
	tests := []struct {
		name, given, expected string
	}{
		{"First-test", "Bye", "Bye"},
		{"Second test", "Hi", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewPrinter(test.given)
			assert.Equal(t, test.expected, got.Value())
		})
	}
}
