package parts

import (
	"fmt"
	"time"
)

func f(n *int) {

	for i := 0; i < 10; i++ {
		fmt.Println(*n, ":", i)
		time.Sleep(time.Second)
	}
}

func Gorutunes() {

	input := 0

	go f(&input)

	fmt.Scanln(&input)

	fmt.Scanln(&input)
}
