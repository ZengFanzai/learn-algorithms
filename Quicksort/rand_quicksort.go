package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Quicksort(array []int, p, r int) {
	if p < r {
		q := RandomPartition(array, p, r)
		Quicksort(array, p, q-1)
		Quicksort(array, q+1, r)

	}
}

func RandomPartition(array []int, p, r int) int {
	i := rand.Intn(r + 1 - p) + p
	array[r], array[i] = array[i], array[r]
	return Partition(array, p, r)
}

func Partition(array []int, p, r int) int {
	x := array[r] //选择主元
	i := p - 1
	for j := p; j < r; j++ {
		if array[j] <= x {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[r] = array[r], array[i+1]
	return i + 1
}

func main() {
	// array := []int{
	// 	55, 94, 87, 1, 4, 2, 32, 11, 77, 222, 39, 42, 64, 53, 70, 12, 9,
	// 	2,8,7,1,3,5,6,4,
	// }
	var array1 []int
	rand.Seed(time.Now().Unix())
	// max := 1000000
	// for index := 0; index < max; index++ {
	// 	array1 = append(array1, rand.Intn(max))
	// }

	//最坏情况
	//两个子数组分别包含n-1个和0个元素
	//100000花费的时间为3.36s左右；1000000花费的时间很长
	max := 1000000
	for index := max-1; index >= 0; index-- {
		array1 = append(array1, index)
	}


	start := time.Now()
	Quicksort(array1, 0, len(array1)-1)
	fmt.Println(time.Now().Sub(start))
	// fmt.Println(array1)

}
