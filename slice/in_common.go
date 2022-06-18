package slice

// InCommon receives two slice of type T, and returns a
// slice with all elements in common between them.
func InCommon[T comparable](slice1, slice2 []T) (common []T) {
	m := map[T]struct{}{}

	for _, v := range slice2 {
		m[v] = struct{}{}
	}

	// 1,3,5,7
	for _, v := range slice1 {
		if _, ok := m[v]; ok {
			common = append(common, v)
			delete(m, v)
		}
	}

	return
}
