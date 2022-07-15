package main

import (
	"fmt"
)

//import "github.com/seweide/goin"

/*
*
占位符	说明
*/
func main() {
	var (
		a = 100
		b = "henry"
	)
	// General（通用占位符）
	//占位符	说明
	//%v	以默认的方式打印变量的值（万能占位符，如果不知道变量是什么类型，用%v即可，go语言会自动为你识别）
	//%T	打印变量的类型
	//%%	字面上的百分号，并非值的占位符

	// %v俗称占位符
	fmt.Printf("a=%v\n", a)
	// %T 打印类型
	fmt.Printf("b的类型是%T\n", b)
	// %%转义
	fmt.Printf("%d%%\n", a)

	//Integer（整型）
	//占位符	说明
	//%+d	带符号的整型
	//%q	打印单引号
	//%o	不带零的八进制
	//%#o	带零的八进制
	//%x	小写的十六进制
	//%X	大写的十六进制
	//%#x	带0x的十六进制
	//%U	打印Unicode字符
	//%#U	打印带字符的Unicode
	//%b	打印整型的二进制

	//fmt.Println("我要拥抱go！")
	//fmt.Println(122)
	//fmt.Println("%a - %b",42,43)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	//Integer width（指定长度的整型，以5为例）
	//占位符	说明
	//%5d	整型长度为5，右对齐，左边留白
	//%-5d	左对齐右边留白
	//%05d	数字前面补零
	fmt.Printf("|%d|\n", a)
	fmt.Printf("|%5d|\n", a)
	fmt.Printf("|%-5d|\n", a)
	fmt.Printf("|%05d|\n", a)

	//Float（浮点数）
	//占位符	说明
	//%f	(=%.6f) 6位小数点
	//%e	(=%.6e) 6位小数点（科学计数法）
	//%g	用最少的数字来表示
	//%.3g	最多3位数字来表示
	//%.3f	最多3位小数来表示
	f1 := 3.141592654
	fmt.Printf("%.2f\n", f1) //最多2位小数表示
	fmt.Printf("%.2g\n", f1) //最多用2位数字表示

}
