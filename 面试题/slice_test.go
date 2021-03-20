package 面试题

import (
	"fmt"
	"testing"
)

func Add(array []int) {
	array = append(array, 1, 2, 3, 4)
}

func Update(array []int) {
	array[0] = 10
}

func TestSlice(t *testing.T) {
	array := []int{0}
	Update(array)

	fmt.Println(array)

	Add(array)
	fmt.Println(array)
}
