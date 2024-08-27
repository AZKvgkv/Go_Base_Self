# Slice、Range and Map

## Slice

Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，所以有了切片以适应动态数组的需求

### 定义切片

声明一个未指定大小的数组来定义切片,切片不需要说明长度

```go
var identifier []type
```

或者可使用`make()`函数来创建切片：

```go
var slice1 []type = make([]type, len)
```

也可以简写为：

```go
slice := make([]type, len)
```

也可以指定容量，其中`capacity`为可选参数：

```go
slice := make([]type, len, capacity)
```

### 切片的初始化

直接初始化切片,`[]`表示切片类型, `{1, 2, 3}`初始值依次为1, 2, 3 其中`capacity = len = 3`

```go
s := [] int {1, 2, 3}
```

初始化切片`s`,是数组`arr`的引用

```go
s := arr[:]
```

### len()和cap()函数

切片是可索引的，并且可以使用`len()`获取长度。
切片提供了计算容量的方法`cap()`可以测量切片最长可以达到多少。

### 切片截取

设置上下限，如`[lower-bound:upper-bound]`

### append() 和 copy()函数

`append()`函数用于在切片末尾添加元素，并返回新的切片。
`copy()`函数用于复制切片，并返回复制的长度。

## Range

`range`关键字用于 for 循环中迭代数组(array)、切片(slice)、字符串(string)或集合(map)的元素。在数组和切片中它返回元素的索引值，在字符串中返回字符的 Unicode 代码。

```go
for i, v := range oldSlice {
    newSlice = append(newSlice, v)
}
```

## Map

Go 语言中的 map 是一种无序的键值对集合。(由于其是哈希表实现的)

```go
var map_variable map[key_data_type]value_data_type

map_variable := make(map[key_data_type]value_data_type)
```

其中`key_data_type`表示键的数据类型，`value_data_type`表示值的数据类型。
如果不初始化`map`,那么就会创建一个`nil map`。而`nil map`不能用来存放键值对。

### delete()函数

用于删除集合的元素，参数为`map`和其对应的`key`。
