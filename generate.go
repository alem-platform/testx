package testx

import (
	"fmt"
	"math/rand"
)

// RandMatrix generates a random slice of slices,
// where n is the size of the outer and m is the size of the inner slices.
func RandMatrix[T any](n, m int, fn func() T) [][]T {
	result := make([][]T, 0, n)

	for i := 0; i < n; i++ {
		result = append(result, RandSlice(m, fn))
	}
	return result
}

// RandSlice generates a random slice,
// where n is the size of the slice.
func RandSlice[T any](n int, fn func() T) []T {
	result := make([]T, 0, n)

	for i := 0; i < n; i++ {
		result = append(result, fn())
	}
	return result
}

func RandInt() int {
	return rand.Intn(300)
}

func RandIntPtr() *int {
	return getPtr(RandInt())
}

func RandRune() rune {
	return rune(rand.Intn(256))
}

func getPtr[T any](v T) *T {
	return &v
}

func RandValidStdin() string {
	return fmt.Sprintf("%d %d", rand.Intn(100), rand.Intn(100))
}

func RandInvalidStdin() string {
	return fmt.Sprintf("%d %c", rand.Intn(100), rune(rand.Intn(100)))
}
