package Problem0334

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{

	{[]int{1, 4, 3, 2, 5}, true},
	{[]int{2, 5, 3, 4, 5}, true},
	{[]int{5, 4, 3, 2, 1}, false},

	// 可以有多个 testcase
}

func Test_increasingTriplet(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, increasingTriplet(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_increasingTriplet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			increasingTriplet(tc.nums)
		}
	}
}