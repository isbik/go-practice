package parts

import "fmt"

func Variables() {

	fmt.Print("Enter a temperature: ")

	var input float32

	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9

	fmt.Println(output)

	fmt.Print("Enter a fut: ")

	fmt.Scanf("%f", &input)

	output = float32(FutToMeter(float64(input)))

	fmt.Println(output)
}

func FutToMeter(value float64) float64 {

	return value * 0.3048
}
