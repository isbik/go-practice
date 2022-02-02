package parts

import (
	"fmt"
	"io/ioutil"
)

func StandardLibrary() {

	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	str := string(bs)

	fmt.Println(str)
}