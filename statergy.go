package main

import "fmt"
import "math/rand"


var unSortedArray [10]int = [10]int{1,2,90,87,8,54,70,32,3,10}
var SortedArray []int


type Sort interface {
	sortType() []int
}
var sortInterface Sort

type bubbleSort struct {

}

func (bs bubbleSort) sortType() []int{

	fmt.Println("\nInside BubbleSort Implementation")
	array := unSortedArray
	for i:=1; i< len(array); i++ {
		for j:=0; j < len(array)-i; j++ {
			if (array[j] > array[j+1]) {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array[:]
}

func (ms mergeSort) sortType() []int{
	fmt.Println("\nInside MergeSort Implementation")
	arr := MergeSortMain(unSortedArray[:])
	return arr[:]
}

type mergeSort struct {

}

func MergeSortMain(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSortMain(slice[:mid]), MergeSortMain(slice[mid:]))
}


type quickSort struct {

}


func (qs quickSort) sortType() []int{
	fmt.Println("\nInside QuickSort Implementation")
	arr := QuickSortMain(unSortedArray[:])
	return arr[:]
}


func QuickSortMain(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickSortMain(less), QuickSortMain(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}


func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}



func doSort() []int {

	return sortInterface.sortType()

}

func changeStatergy(sortInterfaceUser Sort){
	sortInterface = sortInterfaceUser
}

func main()  {

	b := bubbleSort{}
	m := mergeSort{}
	q := quickSort{}

	fmt.Println("Unsorted Array:",unSortedArray)
	changeStatergy(b)
	SortedArray = sortInterface.sortType()
	fmt.Println("Sorted Array using BubbleSort:",SortedArray)

	SortedArray = SortedArray[:0]
	fmt.Println("\nSorted Array after cleaning:",SortedArray)

	changeStatergy(m)
	SortedArray = sortInterface.sortType()
	fmt.Println("Sorted Array using Merge sort:",SortedArray)

	SortedArray = SortedArray[:0]
	fmt.Println("\nSorted Array after cleaning:",SortedArray)

	changeStatergy(q)
	SortedArray = sortInterface.sortType()
	fmt.Println("Sorted Array using Merge sort:",SortedArray)
}
