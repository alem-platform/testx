package testx

import (
	"fmt"
	"math/rand"
)

// NewMatrix generates a random slice of slices,
// where n is the size of the outer and m is the size of the inner slices.
func NewMatrix[T any](n, m int, fn func() T) [][]T {
	result := make([][]T, 0, n)

	for i := 0; i < n; i++ {
		result = append(result, NewSlice(m, fn))
	}
	return result
}

// NewSlice generates a random slice,
// where n is the size of the slice.
func NewSlice[T any](n int, fn Getter[T]) []T {
	result := make([]T, 0, n)

	for i := 0; i < n; i++ {
		result = append(result, fn())
	}
	return result
}

type Getter[T any] func() T

func RandRange(low, high int) Getter[int] {
	return func() int {
		return low + rand.Intn(high-low+1)
	}
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

func RandIntStr() string {
	return fmt.Sprintf("%d", rand.Intn(100))
}

func RandRuneStr() string {
	return fmt.Sprintf("%c", rune(rand.Intn(256)))
}
