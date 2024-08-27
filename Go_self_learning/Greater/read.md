# Reflect and 并发

## 反射

反射是指在运行时动态获取类型信息的能力。通过反射，可以做到在运行时创建对象、调用方法、获取变量的值等。

Go语言的反射包提供了两个主要的功能：

* 反射类型：通过反射可以获取类型信息，包括结构体、接口、函数等。
* 反射值：通过反射可以获取变量的值，包括基本类型、数组、结构体、指针、接口等。

### Typeof

`reflect.Typeof` 函数可以获取变量的类型。

### ValueOf

`reflect.ValueOf` 函数可以获取变量的值。

## Go 语言并发

并发：同一时间段内，多个任务（线程）执行。

并行：同一时刻，多个任务（线程）执行。

Go 语言提供了三个并发机制：

* 并发执行：Go 语言支持并发执行，通过 goroutine 和 channel 实现。
* 同步机制：Go 语言提供了互斥锁、读写锁、条件变量等同步机制。
* 并发调度：Go 语言的运行时调度器负责管理 goroutine 的调度和执行。

### Goroutine

Goroutine 是 Go 语言提供的一种轻量级线程，可以与其他 goroutine 并发执行。
启动单个goroutine

```go
package main

import "fmt"

func hello(){
    fmt.Println("hello golang")
}

func main(){
    go hello()
    fmt.Println("main function")
}
```

上面的代码没有打印 `hello golang` ，因为 `hello` 函数是个 goroutine，它并没有执行，只是启动了一个新的 goroutine。
可以使用 `time.Sleep` 函数来等待 goroutine 执行完毕。

```go
package main

import(
    "fmt"
    "time"
)

func hello(){
    fmt.Println("hello golang")
}

func main(){
    go hello()
    fmt.Println("main function")
    time.Sleep(time.Second)
}
```

sync. WaitGroup(Swg.go)
启动多个goroutine(mul.go)

### Channel

Channel 是 Go 语言提供的用于 goroutine 之间通信的机制。
通道遵循先进先出（FIFO）原则，也就是说，发送方发送的数据，必须先被接收方接收。
每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
`channel` 类型, `channel` 零值为 `nil` 。

```go
var variable_name chan element_type
```

`chan` 关键字用于声明通道。

```go
ch := make(chan int)
```

`make` 函数用于创建通道。

```go
ch := make(chan int)
ch <- 10
```

`ch <-` 表示向通道发送数据。

```go
x := <-ch
```

`<-ch` 表示从通道接收数据。

```go
close(ch)
```

`close` 函数用于关闭通道。

* 无缓冲通道(buffer-free.go)
    只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段

* 有缓冲区的通道(buffer.go)

判断通道关闭

```go
select {
case <-ch:
    fmt.Println("channel is closed")
default:
    fmt.Println("channel is not closed")
}
```

`select` 语句用于判断通道是否关闭。

单向通道

```go
ch := make(chan<- int)
```

`chan<-` 表示只能发送，不能接收。

```go
ch <- 10
```

`<-chan` 表示只能接收，不能发送。

```go
x := <-ch
```

不能使用 `ch <-` 来发送数据。

```go
close(ch)
```

不能使用 `close(ch)` 来关闭单向通道。

### 并发执行

Go 语言支持并发执行，通过 goroutine 和 channel 实现。

* goroutine：Go 语言提供的轻量级线程，可以与其他 goroutine 并发执行。
* channel：Go 语言提供的用于 goroutine 之间通信的机制。

### 并发调度

Go 语言的运行时调度器负责管理 goroutine 的调度和执行。

* 主 goroutine：程序的入口，也是唯一的主 goroutine。
* 系统 goroutine：运行时管理 goroutine 的 goroutine，包括调度器、垃圾回收器等。
* 用户 goroutine：由程序创建的 goroutine。

### 同步机制

Go 语言提供了互斥锁、读写锁、条件变量等同步机制。

### 互斥锁

互斥锁（Mutex）是一种用于保护共享资源的同步机制。

```go
var mu sync.Mutex
```

> 读写互斥锁(RWMutex.go)

读写锁分为两种：读锁（Read Lock）和写锁（Write Lock）。
当一个`goroutine`获取到读锁之后，其他的如果`goroutine`是获取读锁会继续获得读锁，如果是获得写锁就会等待；而当一个`goroutine`获取写锁之后，其他的`goroutine`无论是获得读锁还是写锁都会等待。
