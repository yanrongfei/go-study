package main

import (
	"fmt"
	"strconv"
)

/**
基本数据类型和字符串标识的转换
*/
func main() {
	atoiDemo()
	itoaDemo()
}

func atoiDemo() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("conv :", i1)
	}
}

func itoaDemo() {
	i := 100
	itoa := strconv.Itoa(i)
	fmt.Println("itoa:", itoa)
}
