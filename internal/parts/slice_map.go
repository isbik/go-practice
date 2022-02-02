package parts

import "fmt"

func SliceAndMap() {

	var res [][]int

	for i := 0; i < 5; i++ {
		var row []int

		for j := 0; j < 5; j++ {
			row = append(row, i*j)
		}

		res = append(res, row)
	}

	fmt.Println(res)
}
