package main

import "fmt"

func main() {
	//第一种  定义变量
	var i int
	fmt.Println("i=", i) //int默认值是0
	//赋值
	i = 10
	//使用
	fmt.Println("i=", i)

	var str string
	fmt.Println(str) //string默认是空串
	var a float32
	fmt.Println(a) //float默认值是0

	var j int = 1010 //声明时直接赋值数据类型可以省略
	fmt.Println(j)

	//第二种 根据值进行类型推导
	var num = 1.11
	fmt.Println(num)

	//第三种
	name := "yong"
	//i:=111  错误，这种变量必须是没声明过的
	fmt.Println(name, i)

	//第四种  同时声明多个同类型变量
	var n1, n2, n3, n4 int
	fmt.Println(n1, n2, n3, n4)

	//第五种，同时声明多个多类型变量
	var n5, n6, n7 = 100, "tom", 888
	fmt.Println(n5, n6, n7)

	//第六种 一次性声明多个变量
	a1, a2, a3 := 100, "tom", "888"
	fmt.Println(a1, a2, a3)

	//运算
	fmt.Println(n5 + n7) //数字是加法运算
	//fmt.Println(n5+n6) //不同数据类型不能运算
	fmt.Println(n6 + a2) //字符串是拼接

	//P37
}
