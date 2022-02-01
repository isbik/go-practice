package parts

import (
	"fmt"
	"os"
)

func StandardLibrary() {

	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	stat, err := file.Stat()
	
	if err != nil {
		fmt.Println(err)
	}

	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)

	if err != nil {
		fmt.Println(err)
	}

	str := string(bs)
	fmt.Println(str)
}
