package testx

import (
	"fmt"
	"testing"
)

func Failed(t *testing.T, exp, act, input any) {
	t.Helper()
	fmt.Printf("input: %v\n", input)
	fmt.Printf("actual: %v\n", act)
	fmt.Printf("expected: %v\n", exp)
	t.FailNow()
}
