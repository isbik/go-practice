package parts

import "fmt"

func Functions() {

	defer func() {
		str := recover()
		fmt.Println(str)

		defer first()
		second()
	}()

	panic("Panic")

}

func first() {
	fmt.Println("first")
}
func second() {
	fmt.Println("second")
}
