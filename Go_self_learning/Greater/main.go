package main
import (
	"fmt"
	"reflect"
)
func main() {
	var book_num float32 = 6
	var is_book bool = true
	book_author := "AZ"
	book_detail := make(map[string]string)
	book_detail["author"] = book_author
	fmt.Println(reflect.TypeOf(book_num))
	fmt.Println(reflect.TypeOf(is_book))
	fmt.Println(reflect.TypeOf(book_author))
	fmt.Println(reflect.TypeOf(book_detail))

	fmt.Println(reflect.ValueOf(book_num))
	fmt.Println(reflect.ValueOf(is_book))
	fmt.Println(reflect.ValueOf(book_author))
	fmt.Println(reflect.ValueOf(book_detail))


	address := "AZ"
	// reflect_set_value_one(address)
	// 反射修改值必须通过传递变量地址来修改（float32 就不会报错）。 若函数传递的参数是值拷贝，则会报错。

	reflect_set_value_two(&address)
	fmt.Printf("address: %v\n", address)
}

// 通过反射设置值
// func reflect_set_value_one(x interface{})  {
// 	value := reflect.ValueOf(x)
// 	if value.Kind() == reflect.String {
// 		value.SetString("hello")
// 	}
// }

func reflect_set_value_two(x interface{})  {
	value := reflect.ValueOf(x)
	// 反射中使用 Elem() 方法可以获取指针指向的值
	if value.Elem().Kind() == reflect.String {
        value.Elem().SetString("world")
    }
	
}