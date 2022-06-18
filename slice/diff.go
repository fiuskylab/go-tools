package slice

// Diff receives two slices of type T, and returns a
// slice with all elements not in common between them.
func Diff[T comparable](slice1, slice2 []T) []T {
	diff := []T{}

	m := make(map[T]struct{}, len(slice1))

	for _, i := range slice1 {
		m[i] = struct{}{}
	}

	for _, i := range slice2 {
		if _, ok := m[i]; !ok {
			diff = append(diff, i)
			continue
		}
		delete(m, i)
	}

	for k := range m {
		diff = append(diff, k)
	}

	return diff
}
