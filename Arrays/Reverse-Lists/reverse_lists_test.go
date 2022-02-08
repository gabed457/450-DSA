package reverse_lists

import (
	"fmt"
	"reflect"
	"testing"
)

type reverseListIntTest struct {
	arg1, expected []int
}
type reverseListFloat32Test struct {
	arg1, expected []float32
}

var reverseListIntTests = []reverseListIntTest{
	reverseListIntTest{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	reverseListIntTest{[]int{-1, 2, -3, 4, -5}, []int{-5, 4, -3, 2, -1}},
	reverseListIntTest{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
}
var reverseListFloat32Tests = []reverseListFloat32Test{
	reverseListFloat32Test{[]float32{1, 2, 3, 4, 5}, []float32{5, 4, 3, 2, 1}},
	reverseListFloat32Test{[]float32{-1, 2, -3, 4, -5}, []float32{-5, 4, -3, 2, -1}},
	reverseListFloat32Test{[]float32{1, 2, 3, 4, 5}, []float32{5, 4, 3, 2, 1}},
}

func TestReverseLists(t *testing.T) {
	for _, test := range reverseListIntTests {
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
