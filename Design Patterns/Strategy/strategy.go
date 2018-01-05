package main

import "fmt"
import "math/rand"

//Sorting strategy interface
type SortingStrategy interface{
	sort(data []int) []int
}

type Strategy struct {
	SortingStrategy SortingStrategy
}

//Strategy implements Sorting Strategy
func (o *Strategy) Operate(a []int) []int {
	return o.SortingStrategy.sort(a)
}

//Bubble Sort strategy
type BubbleSort struct{}

func (BubbleSort) sort(x []int) []int {
 	//fmt.Printf("Inside bubble sort...") 
 	end := len(x) - 1
	for {
		if end == 0 {
			break
		}
		for i := 0; i < len(x)-1; i++ {
			if x[i] > x[i+1] {
				x[i], x[i+1] = x[i+1], x[i]
			}
		}
		end -= 1
	}
	//fmt.Println(x)
	return x
}

//Quick Sort strategy
type QuickSort struct{}

func (quicksort QuickSort) sort(slice []int) []int{
 	//fmt.Println("Inside quick sort...") 
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

	less, more = quicksort.sort(less), quicksort.sort(more)

	less = append(less, middle...)
	less = append(less, more...)

	//fmt.Println(less)
	return less
}

//Merge Sort strategy
type MergeSort struct{}

func (mergesort MergeSort) sort(A []int) []int{
	//fmt.Printf("Inside merge sort...") 
    if len(A) <= 1 {
        return A
    }

    left, right := split(A)
    left = mergesort.sort(left)
    right = mergesort.sort(right)
    return merge(left, right)
}

// split array into two
func split(A []int) ([]int, []int) {
    return A[0:len(A) / 2], A[len(A) / 2:]
}

func merge(A, B []int) []int {
    arr := make([]int, len(A) + len(B))
    j, k := 0, 0

    for i := 0; i < len(arr); i++ {
        if j >= len(A) {
            arr[i] = B[k]
            k++
            continue
        } else if k >= len(B) {
            arr[i] = A[j]
            j++
            continue
        }
        if A[j] > B[k] {
            arr[i] = B[k]
            k++
        } else {
            arr[i] = A[j]
            j++
        }
    }

    return arr
}

func runStrategy(){
	x := []int{4,1,6,8,9}
	algo := Strategy{BubbleSort{}}
	x=algo.Operate(x) 
	fmt.Println("Bubble sort result",x)

	//Reset Data
	x = []int{4,1,6,8,9}
	algo = Strategy{QuickSort{}}
	x=algo.Operate(x) 
	fmt.Println("Quick sort result",x)
	
	//Reset Data
	x = []int{4,1,6,8,9}
	algo = Strategy{MergeSort{}}
	x=algo.Operate(x) 
	fmt.Println("Merge sorting result",x)

}

func main(){
	runStrategy()
}
