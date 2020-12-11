package main

import (
	"fmt"
	time "time"
)

func main() {



	timeDemo()
	timestampDemo()
	operator()
	tickDemo()

}

func timeDemo() {
	now := time.Now()
	fmt.Printf("current time %v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	fmt.Printf("%d-%02d-%02d \n", year, month, day)
}

func timestampDemo() {
	now := time.Now()
	unix := now.Unix()
	fmt.Printf("timestamp %v\n", unix)
	//将时间戳转成time
	t := time.Unix(unix, 0)
	fmt.Printf("时间戳转成的时间为%v\n", t)
}

func operator() {
	now := time.Now()
	add := now.Add(time.Hour * 2)
	fmt.Printf("当前时间加2小时：%v\n", add)
	duration, err := time.ParseDuration("-5h")
	t := now.Add(duration)
	fmt.Printf("当前时间加-5小时：%v%v \n", t, err)
	sub := now.Sub(add)
	fmt.Printf("两个时间的间隔%v \n", sub)
}

/**
创建一个定时器
*/
func tickDemo() {
	tick := time.Tick(time.Second * 2)
	for t := range tick {
		fmt.Println(t)
	}
}
