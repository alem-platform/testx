package testx

import (
	"fmt"
	"math/rand"
)

func RandMatrix[T any](n int, fn func() T) [][]T {
	result := make([][]T, 0, n)

	for i := 0; i < n; i++ {
		tmp := []T{}
		for j := 0; j < len(tmp); j++ {
			tmp[j] = fn()
		}
		result = append(result, tmp)
	}
	return result
}

func RandSlice[T any](n int, fn func() T) []*T {
	result := make([]*T, 0, n)

	for i := 0; i < n; i++ {
		result = append(result, getPtr(fn()))
	}
	return result
}

func RandInt() int {
	return rand.Intn(300)
}

func RandRune() rune {
	return rune(rand.Intn(256))
}

func getPtr[T any](v T) *T {
	return &v
}

func RandStdin(n int) []string {
	stdNum := "0123456789"
	resArr := []string{}
	for i := 0; i < 10; i++ {
		res := ""
		j := rand.Intn(9)
		k := rand.Intn(9)

		res = fmt.Sprintf("%v %v", stdNum[j], stdNum[k])
		if i == rand.Intn(9) {
			res = fmt.Sprintf("%v %v", stdNum[j], rune(rand.Intn(256)))
		}
		if i == rand.Intn(9) {
			res = fmt.Sprintf("%v", rune(rand.Intn(256)))
		}
		if i == rand.Intn(9) {
			res = fmt.Sprintf("%v 0", stdNum[j])
		}
		resArr = append(resArr, res)
	}
	return resArr
}
