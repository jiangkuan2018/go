package tempconv

//
//import "fmt"
//
//type Celsius float64
//type Fahrenheit float64
//
//func (c Celsius) String() string {
//	return fmt.Sprintf("%g°C", c)
//}
//
//const (
//	AbsoluteZeroC Celsius = -273.15
//	FreezingC     Celsius = 0
//	BoilingC      Celsius = 100
//)
//
//func CToF(c Celsius) Fahrenheit {
//	return Fahrenheit(c*9/5 + 32)
//}
//func FToC(f Fahrenheit) Celsius {
//	return Celsius((f - 32) * 5 / 9)
//}
//func main() {
//	fmt.Printf("%g\n", BoilingC-FreezingC)       // 摄氏度之间运算
//	boilingF := CToF(BoilingC)                   // 计算华氏度值
//	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // 华氏度计算
//	//fmt.Printf("%g\n", boilingF-FreezingC)       // type error 不允许华氏度和摄氏度运算
//
//	/**
//	比较运算符==和<也可以用来比较一个命名类型的变量和另一个有相同类型的变量，
//	或有着相同底层类型的未命名类型的值之间做比较。
//	但是如果两个值有着不同的类型，则不能直接进行比较：
//	*/
//	var c Celsius
//	var f Fahrenheit
//	fmt.Println(c == 0) // 摄氏度底层类型是64位浮点型
//	fmt.Println(f >= 0) // 华氏度底层类型是64位浮点型
//	//fmt.Println(c == f) // 类型错误
//	fmt.Println(c == Celsius(f)) // 类型转换
//	var res = FToC(boilingF)
//	fmt.Println(res.String())
//	c = FToC(212.0) // 摄氏度类型
//	//	许多类型都会定义一个String方法，因为当使用fmt包的打印方法时，
//	//	将会优先使用该类型对应的String方法返回的结果打印，
//
//	fmt.Println(c.String()) // 100°C
//	fmt.Printf("%v\n", c)   // 100°C
//}
