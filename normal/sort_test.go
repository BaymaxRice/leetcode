package normal

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func TestMergeSort(t *testing.T) {
	data := ints
	//fmt.Printf("merge sort: %+v\n", MergeSort(data[:]))
	fmt.Printf("heap sort: %+v\n", HeapSort(data[:]))
	sort.Ints(data[:])
	fmt.Printf("actual sort: %+v\n", data)
}
