package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	assert.Equal(t, []string{"chirp"}, Keys(map[string]int{"chirp": 47}))
}

func TestMakeRange(t *testing.T) {
	cases := []struct {
		i        int
		j        int
		expected []int
	}{
		{
			i:        43,
			j:        47,
			expected: []int{43, 44, 45, 46},
		},
		{
			i:        47,
			j:        43,
			expected: nil,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("makeRange(%d,%d)", tc.i, tc.j), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					assert.Contains(t, r, "invalid: must have i <= j")
				} else if tc.expected == nil {
					t.Error("expected panic but didn't")
				}
			}()
			assert.Equal(t, tc.expected, MakeRange(tc.i, tc.j))
		})
	}
}
