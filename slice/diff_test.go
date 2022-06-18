package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiff(t *testing.T) {
	t.Run("Diff", func(t *testing.T) {
		require := require.New(t)
		slice1 := []int{1, 2, 3}
		slice2 := []int{4, 5, 6}

		got := Diff[int](slice1, slice2)
		want := append(slice2, slice1...)
		require.ElementsMatch(want, got)
	})

	t.Run("Diff", func(t *testing.T) {
		require := require.New(t)
		slice1 := []int{1, 2, 3}
		slice2 := []int{1, 2, 3, 4}

		got := Diff[int](slice1, slice2)
		want := []int{4}
		require.ElementsMatch(want, got)
	})
}
