package main

import "fmt"

type (
	// UInteger type constraint wraps all unsigned intergers.
	UInteger interface {
		uint | uint64 | uint32 | uint16 | uint8
	}
	// Integer type constraint wraps all intergers.
	Integer interface {
		UInteger |
			int | int64 | int32 | int16 | int8
	}
	// Float type constraint wraps all floats.
	Float interface {
		float64 | float32
	}
	// Number type constraint wraps all number types
	Number interface {
		Integer | Float
	}

	// Iterator type constraint wraps data that can
	// be iterated.
	Iterator[T any] interface {
		[]T | ~map[any]T
	}
)

// Sum receives a
func Sum[U Iterator[T], T Number](vals U) T {
	var res T
	switch vals := any(vals).(type) {
	case []T:
		for _, v := range vals {
			res += v
		}
	case map[any]T:
		for _, v := range vals {
			res += v
		}
	}
	return res
}

func In[U Iterator[T], T Number](item T, vals U) bool {
	switch vals := any(vals).(type) {
	case []T:
		for _, v := range vals {
			if item == v {
				return true
			}
		}
	case map[any]T:
		for _, v := range vals {
			if item == v {
				return true
			}
		}
	}
	return false
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(In(2, slice))
}

func InSlice[U ~[]T, T comparable](item T, slice U) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
