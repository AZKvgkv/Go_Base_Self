package main

import (
	"fmt"
)

func main() {
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 20
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element at index %d is %d\n", j, n[j])
	}

	/*数组 -> 5行2列*/
	a := [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {6, 8}}

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

	// 调用getAverage
	var balance = [5]int{1200, 34, 5, 99, 46}
	/*数组作为参数传递给函数*/
	avg := getAverage(balance[:], 5)
	fmt.Printf("The average is %f\n", avg)

	//指针

	var x = 99
	fmt.Printf("x 的地址是 %x\n", &x)
	var ip *int = &x
	fmt.Printf("ip 的地址是 %x\n", ip)
	fmt.Printf("*ip 指向的值是 %d\n", *ip)

	// nil
	var ptr *int
	fmt.Printf("ptr 的值是 %v\n", ptr)
	fmt.Printf("ptr 的值是 %#v\n", ptr)

	/*go指针数组*/
	// 普通数组
	const MAX int = 5
	array := []int{10, 100, 200, 20, 1}
	for i := 0; i < MAX; i++ {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}

	// 指针数组
	var ptr_array [MAX]*int
	for i := 0; i < MAX; i++ {
		ptr_array[i] = &array[i] /* 整数地址赋值给指针数组 */
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("ptr_array[%d] = %d\n", i, *ptr_array[i])
	}
	// 指向指针的指针
	b := 3000
	var ptr_b *int
	var ptr_ptr **int
	ptr_b = &b
	ptr_ptr = &ptr_b
	fmt.Printf("b 的值是 %d\n", b)
	fmt.Printf("指针变量 *ptr_b 的值是 %d\n", *ptr_b)
	fmt.Printf("指向指针的指针变量 **ptr_ptr 的值是 %d\n", **ptr_ptr)

	// 指针作为函数参数
	var q int = 100
	w := 200
	fmt.Printf("交换前 q 的值是 %d\n", q)
	fmt.Printf("交换前 w 的值是 %d\n", w)

	swap(&q, &w)

	fmt.Printf("交换后 q 的值是 %d\n", q)
	fmt.Printf("交换后 w 的值是 %d\n", w)

	// 访问结构体成员
	main_struct_interview()
	// 结构体指针示例
	main_struct_pointer()
}

func getAverage(arr []int, size int) float32 {
	sum := float32(0)
	for i := 0; i < size; i++ {
		sum += float32(arr[i])
	}
	avg := sum / float32(size)
	return avg
}
func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main_struct_interview() {

	//Book 1 的描述
	var Book1 Books /*声明 Book1 为 Books 类型*/
	Book1.title = "Go 语言"
	Book1.author = "AZ"
	Book1.subject = "Programming"
	Book1.book_id = 101

	//Book 2 的描述
	var Book2 Books /*声明 Book2 为 Books 类型*/
	Book2.title = "Python 语言"
	Book2.author = "AZ"
	Book2.subject = "Programming"
	Book2.book_id = 110

	//Book 3 的描述
	var Book3 Books /*声明 Book3 为 Books 类型*/
	Book3.title = "Rust 语言"
	Book3.author = "AZ"
	Book3.subject = "Programming"
	Book3.book_id = 121

	// 打印Book1的信息
	fmt.Printf("Book1 的标题是 %s\n", Book1.title)
	fmt.Printf("Book1 的作者是 %s\n", Book1.author)
	fmt.Printf("Book1 的科目是 %s\n", Book1.subject)
	fmt.Printf("Book1 的编号是 %d\n", Book1.book_id)

	// 打印Book2的信息
	fmt.Printf("Book2 的标题是 %s\n", Book2.title)
	fmt.Printf("Book2 的作者是 %s\n", Book2.author)
	fmt.Printf("Book2 的科目是 %s\n", Book2.subject)
	fmt.Printf("Book2 的编号是 %d\n", Book2.book_id)

	// 打印Book3的信息
	fmt.Printf("Book3 的标题是 %s\n", Book3.title)
	fmt.Printf("Book3 的作者是 %s\n", Book3.author)
	fmt.Printf("Book3 的科目是 %s\n", Book3.subject)
	fmt.Printf("Book3 的编号是 %d\n", Book3.book_id)
}
func main_struct_pointer() {
	var Book1 Books
	var Book2 Books
	/*Book1 的描述*/
	Book1.title = "Zig 语言"
    Book1.author = "AZ"
    Book1.subject = "Programming"
    Book1.book_id = 211

    /*Book2 的描述*/
    Book2.title = "V 语言"
    Book2.author = "AZ"
    Book2.subject = "Programming"
    Book2.book_id = 201

    /*打印 Book1 的信息*/
	printBook(&Book1)

    /*打印 Book2 的信息*/
    printBook(&Book2)
}
func printBook(book *Books) {
	fmt.Printf("书名：%s\n", book.title)
	fmt.Printf("作者：%s\n", book.author)
	fmt.Printf("科目：%s\n", book.subject)
	fmt.Printf("编号：%d\n", book.book_id)
}