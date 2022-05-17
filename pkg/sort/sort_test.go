package sort

import (
	"fmt"
	"log"
	"testing"
)
import "gotest.tools/v3/assert"

func Test_SimpleInsertSortInt(t *testing.T) {

	arr := []int{5, 2, 4, 6, 1, 3}
	sortedArr := []int{1, 2, 3, 4, 5, 6}
	SimpleInsertionSortInt(arr)
	checkIntArrays(t, arr, sortedArr)

}

func Test_SimpleInsertSortNumeric(t *testing.T) {
	arr := []int{5, 2, 4, 6, 1, 3}
	sortedArr := []int{1, 2, 3, 4, 5, 6}
	SimpleInsertionSortNumeric[int](arr)
	checkIntArrays(t, arr, sortedArr)

}

func Test_RecursiveInsertSort(t *testing.T) {

	arr := []int{5, 2, 4, 6, 1, 3}
	sortedArr := []int{1, 2, 3, 4, 5, 6}
	err := RecursiveInsertionSortInt(arr, len(arr))
	if err != nil {
		t.Errorf("returned error %v", err.Error())
	}
	checkIntArrays(t, arr, sortedArr)

}

func Test_MergeSortedIntSimple(t *testing.T) {

	arr := []int{0, 4, 6, 7, 1, 5, 9}
	sortedArr := []int{0, 1, 4, 5, 6, 7, 9}
	err := MergeSortedInt(arr, 0, 4, 7)
	assert.NilError(t, err)
	checkIntArrays(t, arr, sortedArr)
}

func checkIntArrays(t *testing.T, arr []int, sortedArr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] != sortedArr[i] {
			t.Errorf("incorrect value at index %d, expected %d, got %d",
				i, sortedArr[i], arr[i])
		}
	}
}

func Test_MergeSortedIntOneItemAlreadySorted(t *testing.T) {

	arr := []int{0, 4, 8}
	sortedArr := []int{0, 4, 8}
	err := MergeSortedInt(arr, 0, 1, 3)
	assert.NilError(t, err)
	checkIntArrays(t, arr, sortedArr)
}

func Test_MergeSortedIntOneItemToBeSorted(t *testing.T) {

	arr := []int{0, 4, 2}
	sortedArr := []int{0, 2, 4}
	err := MergeSortedInt(arr, 0, 2, 3)
	assert.NilError(t, err)
	checkIntArrays(t, arr, sortedArr)
}

func Test_MergeSortIntSimple(t *testing.T) {

	arr := []int{0, 4, 6, 7, 1, 5, 9}
	sortedArr := []int{0, 1, 4, 5, 6, 7, 9}
	err := MergeSortInt(arr, 0, 6)
	assert.NilError(t, err)
	checkIntArrays(t, arr, sortedArr)
}

func Test_MergeSortIntBookSample(t *testing.T) {

	arr := []int{3, 41, 52, 26, 38, 57, 9, 49}
	sortedArr := []int{3, 9, 26, 38, 41, 49, 52, 57}
	err := MergeSortInt(arr, 0, len(arr))
	assert.NilError(t, err)
	checkIntArrays(t, arr, sortedArr)
}

func Test_MergeWithInsertionSortIntBookSample(t *testing.T) {

	arr := []int{3, 41, 52, 26, 38, 57, 9, 49}
	sortedArr := []int{3, 9, 26, 38, 41, 49, 52, 57}
	err := MergeWithInsertionSortInt(arr, 0, len(arr), 4)
	assert.NilError(t, err)
	checkIntArrays(t, arr, sortedArr)
}

func BenchmarkMergeWithInsertionSortInt(b *testing.B) {
	var arr []int
	maxThreshold := 5
	for i := 0; i < b.N; i++ {
		arr = []int{3, 41, 52, 26, 38, 57, 9, 49}
		threshold := i % maxThreshold
		err := MergeWithInsertionSortInt(arr, 0, len(arr), threshold)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error in benchmark for threshold %d", threshold))
		}
	}

}
