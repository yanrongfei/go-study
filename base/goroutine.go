package main

import (
	"fmt"
	"time"
)

/**
使用go关键字开启一个goroutine
*/
func main() {
	// goroutine 测试
	//go say("hello")
	//say("world")

	// channel 测试
	//s := []int{7, 2, 8, -9, 4, 0}
	//c := make(chan int)
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // 从通道 c 中接收
	//fmt.Println(x, y, x+y)

	// channel 测试
	// make 设置channel的容量 100，若是不设置则没有缓冲器，会阻塞
	c := make(chan int, 100)
	go send(c)
	go receive(c)

	for true {}
}

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

/**
channel 通道
ch <- v //把v发送到ch
v := <-ch //从ch接收数据 并把值赋给v

*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func send(c chan int) {
	i := 0
	for true {
		i++
		time.Sleep(1 * time.Second)
		fmt.Printf("send msg:%d\n", i)
		c <- i
	}
}

func receive(c chan int) {
	for i := range c {
		fmt.Printf("接收到的msg:%d\n", i)
	}

}
