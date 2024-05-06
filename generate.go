package testx

import (
	"fmt"
	"math/rand"
)

func RandArrayInt(n int) [][20]int {
	result := make([][20]int, 0, n)

	for i := 0; i < n; i++ {
		tmp := [20]int{}
		for j := 0; j < len(tmp); j++ {
			tmp[j] = rand.Intn(300)
		}
		result = append(result, tmp)
	}
	return result
}

func RandArrayRune(n int) [][20]rune {
	result := make([][20]rune, 0, n)
	for i := 0; i < n; i++ {
		tmp := [20]rune{}
		randLng := rand.Intn(len(tmp))
		for j := 0; j < len(tmp); j++ {
			if j == randLng {
				continue
			}
			tmp[j] = rune(rand.Intn(256))
		}
		result = append(result, tmp)
	}
	return result
}

func RandIntPtr(n int) []*int {
	result := make([]*int, 0, n)

	for i := 0; i < n; i++ {
		result = append(result, getPtr(rand.Intn(100)))
	}
	return result
}

func RandSliceInt(n int) [][2]int {
	result := make([][2]int, 0, n)

	for i := 0; i < n; i++ {
		tmp := [2]int{}
		for j := 0; j < len(tmp); j++ {
			tmp[j] = rand.Intn(100)
		}
		result = append(result, tmp)
	}
	return result
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
		if i%n == 1 {
			res = fmt.Sprintf("%v %v", stdNum[j], string(rand.Intn(256)))
		}
		if i%n == 3 {
			res = fmt.Sprintf("%v", string(rand.Intn(256)))
		}
		if i%n == 4 {
			res = fmt.Sprintf("%v 0", stdNum[j])
		}
		resArr = append(resArr, res)
	}
	return resArr
}
