package practice

import (
	"fmt"
	"testing"
)

func TestCong(t *testing.T) {
	result := add(3, 4)
	if result != 7 {
		t.Errorf("Kỳ vọng 7, nhưng nhận được %d", result)
	} else {
		fmt.Println("TestCong passed")
	}

	cases := []struct {
		a, b, expect int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, c := range cases {
		got := add(c.a, c.b)
		if got != c.expect {
			t.Errorf("Cong(%d, %d) = %d; muốn %d", c.a, c.b, got, c.expect)
		} else {
			fmt.Printf("TestCong passed for %d + %d = %d\n", c.a, c.b, got)
		}
	}
}
