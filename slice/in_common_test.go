package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInCommon(t *testing.T) {
	t.Run("InCommon", func(t *testing.T) {
		require := require.New(t)

		s1 := []int{1, 3, 5, 7}
		s2 := []int{2, 3, 6, 7}

		want := []int{3, 7}
		got := InCommon[int](s1, s2)

		require.ElementsMatch(want, got)
	})

	t.Run("InCommon", func(t *testing.T) {
		require := require.New(t)

		s1 := []int{1, 3, 5, 7}
		s2 := []int{2, 4, 6, 8}

		want := []int{}
		got := InCommon[int](s1, s2)

		require.ElementsMatch(want, got)
	})

}
