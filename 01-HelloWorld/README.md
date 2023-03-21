## 01-HelloWorld 笔记

操作系统：Windows11 21H2

Go版本： go 1.20

### Go中的“包”

```go
package main
```

- 使用 `package` 声明
- 允许下划线
- 可以和所属文件夹名字不同，但该文件夹下所有go文件中的包声明需一致

```go
// 单个引用
import "fmt"
// 批量引用
import (
    "strings"
    _ "bytes" // 匿名引用
)
```

- 使用`import` 引用
- 允许批量引用
- 允许匿名引用

> 如果下文不使用这个包提供的函数(方法)，那么需要删除这个包引用，否则会报错。但如果需要用到这个包的 init 函数而不使用其他函数的话，那么可以用 `_` 方式匿名引用



### 常用数据类型

- **string**  字符串, 包装类： **strings**
- **int, int8, int16, int32, int64** 带符号整数
- **uint, uint8, uint16, uint32, uint64** 无符号整数
- **float32, float64** 浮点数
- **bool** 布尔型
- **byte = uint8**,  字节，包装类：**bytes**
- **rune = int32**, 可以理解为1个字符

```go
var str1 = "你好，世界！"
fmt.Printf("%s size = %d, len = %d\r\n", str1, len(str1), utf8.RuneCountInString(str1))
```

输出：

```
你好，世界！ size = 18, len = 6
```

- 可以使用 `len(xxx)`方法获取字符串**字节长度**。

- 用 `codeType.RuneCountInString(xxx)`可以获取字符串的**字符数量**。



### 变量声明

```go
// 单个变量声明
var a int = 1 //显式类型声明
var b = 2     //隐式类型声明 类型为int
var c = "hi"  //隐式类型声明 类型为string
var d = 1.2   //隐式类型声明 类型为float64

// 批量变量声明
var (
	aa = 1
    bb = "Hello"
)

func foo(){
    // 局部变量声明
    cc := "Hi"
    println(cc)
}

```

- 变量的首字母控制访问性：首字母大写包外可访问（如果是全局变量）
- Go支持类型推断
- 使用驼峰命名法
- 如果变量声明了但没有使用，则编译不过，会报错
- 在同一个作用域下，变量只能声明一次。
- 常量的声明使用 `const`

>如果在函数外部声明了一个变量a, 函数内部也声明了变量a,  那么在函数内部作用域下，其内部定义的a有效。 



### 函数声明

```go
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
```

- 返回值允许有多个
- 首字母大写包外可访问，否则只允许包内访问

