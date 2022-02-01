package parts

import (
	"fmt"
	"time"
)

func Gorutunes() {

	fmt.Println("tets")
	Sleep(3)
	fmt.Println("tets 2")
}

func Sleep(v int) {
	<-time.After(time.Second * 2)
}
