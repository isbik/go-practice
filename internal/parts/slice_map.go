package parts

import (
	"fmt"
)

func fibonacci() func() int {
	n := 0
	n1 := 1

	return func() int {
		defer func() {
			n, n1 = n1, n+n1
		}()

		return n
	}
}

func SliceAndMap() {

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
