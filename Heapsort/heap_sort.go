package main

import (
	"fmt"
	"math/rand"
	"time"
)

//HeapSort 堆排序
//时间复杂度为O(nlgn)
func HeapSort(array []int) {
	BuildMaxHeap(array)
	for i := len(array); i > 0; i-- {
		array[i-1], array[0] = array[0], array[i-1] //根节点为最大值，与最后一个元素置换
		heapSize := i - 1 //现在只需要对前i-1个元素重新提取最大堆元素
		MaxHeapifyWithForLoop(array, 0, heapSize) //继续维护最大堆性质
	}
}

//MaxHeapify 维护最大堆性质, 保证array[i]的值在最大堆中逐级下降
func MaxHeapify(array []int, i, heapSize int) {
	//左孩子
	l := 2*i + 1
	//右孩子
	r := 2*i + 2
	largest := i
	if l < heapSize && array[l] > array[i] {
		largest = l
	}
	if r < heapSize && array[r] > array[largest] {
		largest = r
	}
	if largest != i {
		array[largest], array[i] = array[i], array[largest]
		MaxHeapify(array, largest, heapSize)
	}
}

//MaxHeapifyWithForLoop 使用for循环改写递归
func MaxHeapifyWithForLoop(array []int, i, heapSize int) {
	for {
		//左孩子
		l := 2*i + 1
		//右孩子
		r := 2*i + 2
		largest := i
		if l < heapSize && array[l] > array[i] {
			largest = l
		}
		if r < heapSize && array[r] > array[largest] {
			largest = r
		}
		if largest == i {
			return
		}
		array[largest], array[i] = array[i], array[largest]
		i = largest
	}
}

//BuildMaxHeap 构建最大堆
func BuildMaxHeap(array []int) {
	heapSize := len(array)
	parent := heapSize / 2 //相当于最后一个元素的父节点
	for i := parent; i >= 0; i-- {
		MaxHeapifyWithForLoop(array, i, heapSize)
	}
}

func main() {
	var array1 []int
	// array := []int{
	// 	55, 94, 87, 1, 4, 2, 32, 11, 77, 222, 39, 42, 64, 53, 70, 12, 9,
	// }
	rand.Seed(time.Now().Unix())
	for index := 0; index < 1000000; index++ {
		array1 = append(array1, rand.Intn(1000000))
	}
	// start := time.Now()
	// fmt.Println(time.Now().Sub(start))
	HeapSort(array1)
	fmt.Println(array1)
}
