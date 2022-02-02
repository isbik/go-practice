package parts

import (
	"fmt"
	"strings"
)

func SliceAndMap() {

	var s string = "Test q q q test test word"

	words := strings.Split(s, " ")

	result := make(map[string]int)

	for _, word := range words {
		result[word] += 1
	}

	fmt.Println(result)
}
