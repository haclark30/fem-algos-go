package quicksort

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	arr := []int{8, 3, 7, 4, 69, 420, 42}

	QuickSort(arr)
	if !reflect.DeepEqual(arr, []int{3, 4, 7, 8, 42, 69, 420}) {
		t.Fatalf("not equal to expected value")
	}
}
