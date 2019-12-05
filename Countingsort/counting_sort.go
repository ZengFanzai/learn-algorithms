package main


import (
	"fmt"
)

func Countingsort(arrayA, arrayB []int, k int)  {
	arrayC := make([]int, k)
	for i := 0; i < k; i++ {
		arrayC[i] = 0
	}
	for j := 0; j < len(arrayA); j++ {
		arrayC[arrayA[j]] = arrayC[arrayA[j]] + 1
	}

	for i := 1; i < k; i++ {
		arrayC[i] = arrayC[i] + arrayC[i-1]
	}

	for j := len(arrayA)-1; j >= 0; j-- {
		arrayB[arrayC[arrayA[j]] - 1] = arrayA[j]
		arrayC[arrayA[j]] = arrayC[arrayA[j]] - 1
	}
}


func main() {
	arrayA := []int{1,4,5,7,1,3,4,10,2,4}
	arrayB := make([]int, len(arrayA))
	Countingsort(arrayA, arrayB, 20)
	fmt.Println(arrayB)
}