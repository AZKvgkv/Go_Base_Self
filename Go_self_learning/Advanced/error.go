package main

import (
	"fmt"
)


// 定义一个 DivideError 结构
type DivideError struct {
	Numerator   int
	Denominator int
}
// 实现 error 接口
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the Denominator is zero.
	Numerator: %d
	Denominator: 0
`
	return fmt.Sprintf(strFormat, de.Numerator)
}

// 定义 int 类型除法运算的函数
func Divide (varNumerator int, varDenominator int) (result int, errorMsg string) {
	if varDenominator == 0{
		dData:=DivideError{
			Numerator: varNumerator,
			Denominator:varDenominator,
			}
			errorMsg = dData.Error()
			return
	}else{
		return varNumerator/varDenominator, ""
	}

}

// func main() {
// 	// 正常情况
// 	if result, errorMsg := Divide(10, 2); errorMsg == "" {
// 		fmt.Printf("10 / 2 = %d\n", result)
// 	}
// 	// 异常情况
// 	if _,errorMsg := Divide(10, 0); errorMsg != "" {
// 		fmt.Println("errorMsg is: ", errorMsg)
// 	}
// }
