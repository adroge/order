package order

// Compare can be used to compare built-in and derived types.
func Compare[
	T ~string |
		~float32 | ~float64 |
		~int | ~int8 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint32 | ~uint64 |
		~uintptr](left, right T) int {
	if left < right {
		return -1
	}
	if left > right {
		return 1
	}
	return 0
}

// ComparePv can be used to compare build-in and derived type values
// referenced by a pointer.
func ComparePv[
	T ~string |
		~float32 | ~float64 |
		~int | ~int8 | ~int32 | ~int64 |
		~uint | ~uint32 | ~uint64 |
		~uintptr](left, right *T) int {
	if *left < *right {
		return -1
	}
	if *left > *right {
		return 1
	}
	return 0
}
