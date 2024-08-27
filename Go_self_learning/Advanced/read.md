# 高级部分

## Go 语言递归函数

语法格式如下：

```go
func recursion(){
    recursion()
}

func main(){
    recursion()
}
```

* 阶乘
* 斐波那契数列

## 类型转换

```go
type_name(expression)
```

`type_name` 为类型, `expression` 为表达式。

* go 不支持隐式转换类型

此时会报错

```go
package main

import "fmt"

func main() {
    a := 3
    var b int32
    b = a
    fmt.Printf("b = %d\n", b)
}
```

但是如果改成`b = int32(a)`就不会报错了。

```go
package main

import "fmt"

func main() {
    a := 3
    var b int32
    b = int32(a)
    fmt.Printf("b = %d\n", b)
}
```

## Go 语言接口(api_call.go)

Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法，就可以作为接口来使用。

```go
/*定义接口*/
type interface_name interface {
    method_name1 [return_type]
    method_name2 [return_type]
   ...
    method_name_n [return_type]
}

/*定义结构体*/
type struct_name struct {
    // variables
}

/*实现接口方法*/
func (struct_name_variable struct_name)method_name1()[return__type]{
    // 方法实现
}
...
func (struct_name_variable struct_name)method_name_n()[return__type]{
    // 方法实现
}
```

## Go 错误处理(error.go)

Go 语言提供了两种错误处理方式：

1. 错误检查：通过检查函数返回的错误值来判断是否有错误发生。
2. 错误恢复：通过 defer 关键字来延迟函数的执行，直到函数返回错误时才会执行。

`error`类型是一个接口类型，这是它的定义：

```go
type error interface{
    Error() string
}
```

可以在编码中通过实现`error`接口类型来生成错误信息。
函数通常在最后的返回值中返回错误信息。使用`errors.New()`可返回一个错误信息。

```go
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
}
```

下面调用代码的实例

```go
result, err := Sqrt(-1)
if err != nil {
    fmt.Println(err)
}
```
