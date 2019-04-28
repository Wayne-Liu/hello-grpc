package main

import (
	"fmt"
	"os"
)

func main()  {
	//os.Getpid return current pid
	str := fmt.Sprintf("%d", os.Getpid())
	bts := []byte(str)

	fmt.Println(str)
	fmt.Println(bts[3])
}
