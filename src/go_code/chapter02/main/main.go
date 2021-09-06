package main

import "fmt"

// 演示转义字符
func main() {
	fmt.Println("tom\tjack")
	fmt.Println("tom\njack")

	fmt.Println("E:\\Projects\\Go\\GoLearn\\src\\go_code\\chapter02\\main\\main.go")
	fmt.Println("tom say\"jack i love you\"")
	fmt.Println("天龙八部雪山飞狐\r张飞厉害") //从当前行的最前面开始输出，覆盖掉以前的内容

}
