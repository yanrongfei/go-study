package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("test.md")
	if err != nil {
		fmt.Println("open file err:", err)
	}
	//延迟调用 在方法return前调用  多个defer先进后出
	//关闭文件
	defer file.Close()

	//1.使用 read方式 只能读取设置好的长度
	tmp := make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))

	//2.循环读取
	var content []byte
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))

	//3.bufio读取文件
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}

	//4.ioutil
	//写入文件
	str := "hello go"
	err = ioutil.WriteFile("test.md", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
	//读取文件
	text, err := ioutil.ReadFile("base/test.md")
	if err != nil {
		fmt.Println("read file failed, err:", err)
	}
	fmt.Println(string(text))



}
