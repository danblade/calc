package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEval(t *testing.T) {

	for _, tt := range []struct {
		input  string
		output int
	}{
		{"1 + 1", 2},
		{"2 + 2", 4},
		{"2 - 8", -6},
		{"16 / 4", 4},
	} {
		t.Run(tt.input, func(t *testing.T) {
			n, err := eval(tt.input)
			require.NoError(t, err)
			require.EqualValues(t, tt.output, n)
		})
	}
}
