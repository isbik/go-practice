package parts

import "fmt"

func Functions() {

	fmt.Println(half(3))
	fmt.Println(half(5))
	fmt.Println(half(6))

	fmt.Println("------------------")

	fmt.Println(max(1, 2, 2, 7, 2, 2))

	fmt.Println("------------------")

	g := makeOddGenerator()

	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

	fmt.Println("------------------")

	fmt.Println(fib(4))
	fmt.Println(fib(9))
	fmt.Println(fib(2))
	fmt.Println(fib(1))

}

func half(num int) (int, bool) {

	return num / 2, num%2 == 0
}

func max(nums ...int) int {
	res := nums[0]

	for _, v := range nums {

		if v > res {
			res = v
		}
	}

	return res
}

func makeOddGenerator() func() int {
	v := 1

	return func() (ret int) {
		ret = v
		v += 2
		return
	}
}

func fib(v int) int {

	if v <= 1 {
		return 0
	}

	if v <= 2 {
		return 1
	}

	return fib(v-2) + fib(v-1)

}
