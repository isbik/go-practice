package parts

import "fmt"

func Pointers() {

	x := 5
	y := 55

	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(x *int, y *int) {
	temp := new(int)

	*temp = *x
	*x = *y
	*y = *temp

	*y, *x = *x, *y
}
