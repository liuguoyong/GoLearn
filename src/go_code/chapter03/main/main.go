package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//第一种  定义变量
	var i int
	fmt.Println("i=", i) //int默认值是0
	//赋值
	i = 10
	//使用
	fmt.Println("i=", i)

	var str string
	fmt.Println(str) //string默认是空串  ""
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
	var n1, n2, n3, n4 int64
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
	//基本数值类型
	/*
				数值整数型：int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,byte(存单个字母字符)，rune
				字节数    4/8    1    2     4    8    4/8    1	  2	      4	    8
				数值浮点型：float32,float64
			      字节数      4       8
				布尔型：true,false
		        字节数   1    1
				字符串：string
				派生数据类型：指针，数组，结构体，管道，函数，切片，接口，map

				保小不保大，尽量使用占用空间小的数据类型
	*/

	fmt.Printf("%T\n", a1) //格式化输出
	fmt.Printf("%T  %d\n", n6, unsafe.Sizeof(n6))
	fmt.Println(unsafe.Sizeof(n1))
	//var a4 byte = '背'  //go语言存储中文字符是utf-8
	var a4 int = '背'
	fmt.Printf("ca4=%c  a4=%d\n", a4, a4)

	var b = false //默认值是false
	fmt.Println(b)
	fmt.Println(unsafe.Sizeof(b))
	var sss = "defsv\nfsdvs"
	//sss[0] = a// 字符串一经赋值就不可变
	fmt.Println(sss)

	str3 := `package main

import (
	"fmt"
	"unsafe"
)

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
	var n1, n2, n3, n4 int64
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
	//基本数值类型
	/*
			数值整数型：int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,byte(存单个字母字符)，rune
			字节数    4/8    1    2     4    8    4/8    1	  2	      4	    8
			数值浮点型：float32,float64
		      字节数      4       8
			布尔型：true,false
	        字节数   1    1
			字符串：string
			派生数据类型：指针，数组，结构体，管道，函数，切片，接口，map

			保小不保大，尽量使用占用空间小的数据类型
	*/

	fmt.Printf("%T\n", a1)//格式化输出
	fmt.Printf("%T  %d\n", n6, unsafe.Sizeof(n6))
	fmt.Println(unsafe.Sizeof(n1))
	//var a4 byte = '背'  //go语言存储中文字符是utf-8
	var a4 int = '背'
	fmt.Printf("ca4=%c  a4=%d\n",a4,a4)

	var b = false
	fmt.Println(b)
	fmt.Println(unsafe.Sizeof(b))
	var sss = "defsv\nfsdvs"
	//sss[0] = a// 字符串一经赋值就不可变
	fmt.Println(sss)


}
`
	fmt.Println(str3) //反引号可输出转义字符
	//字符串拼接方式
	var str4 = "aaa" + "bbb"
	str4 += "ccc"
	fmt.Println(str4)

	var str5 = "111" + "2222" + //加号留在上一行
		"55555" +
		"7777"
	fmt.Println(str5)
}
