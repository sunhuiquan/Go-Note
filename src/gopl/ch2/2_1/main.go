package main

import (
	"fmt"
	"gopl/ch2/2_1/tempconv"
)

func main() {
	k := tempconv.CToK(tempconv.Celsius(0))
	fmt.Println(k)
}
