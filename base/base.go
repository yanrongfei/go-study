package main

import "fmt"

func main() {
	//test00()
	//test01()
	//test02()
	//test03()
	//test04()
	//test05()
	//test06()
	//test07()
	test08()

}

func test00(){
	var name string
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Printf("输入的name：%s \n", name)
}

/**
1.变量定义  数字、字母、下划线 必须以字母下划线开头
2. %d 数字类型  %s 字符串类型

*/

func test01() {
	//变量的声明
	var b, c int = 1, 2
	var (
		a int    = 1
		d int    = 2
		e bool   = false
		f string = "hello"
	)
	fmt.Println(a, b, c, d, e)
	fmt.Printf(f+":%s\n", "zs")
	fmt.Printf("%T\n", a)  //查看数据的类型
	fmt.Printf("%x\n", &a) //16进制 内存地址

	//常量的声明
	const name string = "zs"
	fmt.Print(name)

	// :=  简短变量声明
	g := "ss"
	fmt.Print(g)

}

func test02() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	//定义数量级
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
	)
	fmt.Println(KB, MB)
}

/**
条件判断
*/
func test03() {
	a := 100
	if a <= 10 {
		fmt.Println("a <= 10")
	} else if a <= 50 {
		fmt.Println("a <= 50")
	} else {
		fmt.Println("a > 50")
	}

	// switch
	n := 5
	switch n {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(n)
	}
}

/**
循环
*/
func test04() {

	for a := 0; a < 100; a++ {
		fmt.Printf("循环次数：%d\n", a)
		//当a=50跳出
		if a == 50 {
			break
		}
		//当 a = 80 跳过
		if a == 80 {
			continue
		}

	}

}

/**
指针
*/
func test05() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */
	ip = &a        /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// 空指针
	var ptr *int
	fmt.Printf("%x\n", ptr)
	b := ptr == nil
	fmt.Println(b)

}

func test06() {
	scoreMap := make(map[string]int)
	scoreMap["zs"] = 100
	scoreMap["ls"] = 90
	i, ok := scoreMap["ww"]
	if ok {
		fmt.Println(i)
	} else {
		fmt.Println("查无此人")
	}

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

/**
定义结构体
 */
func test07() {

	type User struct {
		name string
		gender string
		mobile string
		birth int
	}

	user := User{
		name:   "zs",
		gender: "male",
		mobile: "123456",
	}
	fmt.Println(user)

	var user01 User
	user01.name = "ls"
	user01.mobile = "123"
	fmt.Println(user01)

}

func test08()  {
	// <<  3 * 2²
	var a = 3 << 2
	fmt.Println(a)
	// >> 300 / 2²
	 b := 200 >> 2
	fmt.Println(b)

}
