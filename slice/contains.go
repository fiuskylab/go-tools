package slice

// Contains will search in a slice if contains given item.
func Contains[T comparable](item T, slice []T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
