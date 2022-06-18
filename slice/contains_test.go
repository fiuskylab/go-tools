package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	t.Run("Contains - []int - Found", func(t *testing.T) {
		require := require.New(t)

		slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := true
		got := Contains[int](2, slc)

		require.Equal(got, want)
	})

	t.Run("Contains - []int - Not Found", func(t *testing.T) {
		require := require.New(t)

		slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := false
		got := Contains[int](23, slc)

		require.Equal(got, want)
	})
}
