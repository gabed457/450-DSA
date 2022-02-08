package max_min_from_list

import (
	"fmt"
	"reflect"
	"testing"
)

type ListTest struct {
	arg1, expected []int
}

var ListTests = []ListTest{
	ListTest{[]int{1, 2, 3, 4, 5}, []int{1, 5}},
	ListTest{[]int{-1, 2, -3, 4, -5}, []int{-5, 4}},
}

func TestReverseLists(t *testing.T) {
	for _, test := range ListTests {
		if output := maxMinArr(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkReverseLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxMinArr([]int{1, 2, 3, 4, 5})
	}
}

func ExampleMaxMinArr() {
	fmt.Println(maxMinArr([]int{7, 8, 9}))
	//Output: [7 9]
}
