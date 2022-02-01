package parts

import "fmt"

func SliceAndMap() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := make([]int, 4, 5)

	copy(s2, s1)

	fmt.Println(s1, s2)

	x := [6]string{"a", "b", "c", "d", "e", "f"}

	fmt.Println(x[2:5]) //c d e

	nums := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println(FindSmallest(nums))

}

func FindSmallest(nums []int) int {

	smallest := nums[0]

	for _, num := range nums {

		if num < smallest {
			smallest = num
		}
	}

	return smallest
}
