package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	assert.Equal(t, []string{"chirp"}, Keys(map[string]int{"chirp": 47}))
}
