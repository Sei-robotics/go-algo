package sort

import (
	"errors"
	"fmt"
)

type Lessable interface {
	LessThan(thanThis Lessable) bool
}

func SimpleInsertionSort[T Lessable](arr []T) {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && val.LessThan(arr[j]) {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = val
	}
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func SimpleInsertionSortNumeric[T Numeric](arr []T) {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = val
	}
}

func SimpleInsertionSortInt(arr []int) {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = val
	}
}

func RecursiveInsertionSortInt(arr []int, pos int) error {
	if pos > len(arr) {
		return fmt.Errorf("passed size %d greater than array length %d", pos, len(arr))
	}

	if pos == 1 {
		return nil
	}
	if err := RecursiveInsertionSortInt(arr, pos-1); err != nil {
		return err
	}
	i := pos - 1
	val := arr[i]
	j := i - 1
	for j >= 0 && arr[j] > val {
		arr[j+1] = arr[j]
		j -= 1
	}
	arr[j+1] = val
	return nil
}

// merges two sorted int slices
// merge sorted arrays:
// {arr[start], ..., arr[middle-1]} and
// {arr[middle], ..., arr[end-1]}
func MergeSortedInt(arr []int, start int, middle int, end int) error {
	if middle < start || end < middle {
		return errors.New("wrong indices")
	}

	len1 := middle - start
	len2 := end - middle

	left := make([]int, len1)
	right := make([]int, len2)

	for i := 0; i < len1; i++ {
		left[i] = arr[start+i]
	}
	for i := 0; i < len2; i++ {
		right[i] = arr[middle+i]
	}
	i := 0
	j := 0
	for k := start; k < end; k++ {
		if i >= len1 {
			arr[k] = right[j]
			j++
			continue
		}
		if j >= len2 {
			arr[k] = left[i]
			i++
			continue
		}
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	return nil
}

func MergeSortInt(arr []int, start int, end int) error {
	if start+1 >= end {
		return nil
	}
	var middle int = (end + start) / 2
	var err error
	err = MergeSortInt(arr, start, middle)
	if err != nil {
		return err
	}
	err = MergeSortInt(arr, middle, end)
	if err != nil {
		return err
	}
	MergeSortedInt(arr, start, middle, end)
	return nil
}

func MergeWithInsertionSortInt(arr []int, start int, end int, threshold int) error {
	if start+1 >= end {
		return nil
	}
	if end-start <= threshold {
		SimpleInsertionSortInt(arr[start:end])
		return nil
	}
	var middle int = (end + start) / 2
	var err error
	err = MergeWithInsertionSortInt(arr, start, middle, threshold)
	if err != nil {
		return err
	}
	err = MergeWithInsertionSortInt(arr, middle, end, threshold)
	if err != nil {
		return err
	}
	MergeSortedInt(arr, start, middle, end)
	return nil
}
