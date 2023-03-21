package main

import "fmt"

func main() {
	// ---- print title -------
	var str = "Hello World" // 变量定义方式1
	strLen := len(str)      // 变量定义方式2 这种只能局部使用
	// print 和 println 并不支持格式化输出
	print(str + " => ")
	println("length: ", strLen)
	// 如果需要格式化输出，要使用 fmt.Printf
	fmt.Printf("%s => length: %d\r\n", str+"!", strLen+1)

	// ---- print myself information ----
	name, age := GetMySelfInformation()
	infoStr := introduceMyself(name, age)
	println(infoStr)
}

// introduceMyself
// demo: 带参函数，并且具有一个返回值，此函数只有本文件内可用
func introduceMyself(name string, age int) string {
	res := fmt.Sprintf("My name is %s, ", name)
	res += fmt.Sprintf("and I am %d years old.", age)
	return res
}

// GetMySelfInformation 获取个人信息，返回名字和年龄
// 不带参函数，具有两个返回值，此函数全局可用
func GetMySelfInformation() (string, int) {
	return "Jelin", 25
}
