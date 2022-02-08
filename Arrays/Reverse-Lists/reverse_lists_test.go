package reverse_lists

import (
	"fmt"
	"reflect"
	"testing"
)

type reverseListTest struct {
	arg1, expected []int
}

var reverseListTests = []reverseListTest{
	reverseListTest{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	reverseListTest{[]int{-1, 2, -3, 4, -5}, []int{-5, 4, -3, 2, -1}},
	reverseListTest{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
}

func TestReverseLists(t *testing.T) {
	for _, test := range reverseListTests {
		if output := reverseList(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkReverseLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseList([]int{1, 2, 3, 4, 5})
	}
}

func ExampleReverseList() {
	fmt.Println(reverseList([]int{7, 8, 9}))
	//Output: [9 8 7]
}
