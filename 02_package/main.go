package main

import (
	"fmt"
	_ "fmt"
	winniepooh "go_demo/02_package/icomefromalaska"
	"go_demo/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.MyName)
	fmt.Println(stringutil.Reverse("今天是星期五"))
	fmt.Println(winniepooh.BearName)
}
