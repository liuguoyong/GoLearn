package main

import "fmt"

func main() {
	fmt.Println("姓名\t\t\t年龄\t\t\t籍贯\t\t\t住址\njohn\t\t12\t\t\t河北\t\t\t北京")
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
	var num = 5*5 + 3
	fmt.Println(num)
	//超长输出，每行尽量不要超过80字符
	fmt.Println("dgbaihdihwdlhseljkhsdjklfhsdjklhfjalsdhfjashfjda",
		"shjfhasjkhfaklshdfalskhfiuwehfiuwhgferiohguerghuieoshgupeshgue",
		"shrgpehsrguherspghesrphgusehrg")
}
