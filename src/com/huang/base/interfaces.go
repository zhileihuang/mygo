package main

import "fmt"

/**
* 定义一个结构体Square 和一个接口Shaper,接口有一个方法Area()
* 在main()方法中创建了一个Square的实例.在主程序外边定义了一个接收者类型是Square方法的Area(),用来计算正方形的面积:结构体Square实现了接口Shaper
* 所以,可以将一个Square类型的变量赋值给一个接口类型的变量:areaIntf=sq1
**/
type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	areaIntf := sq1
	fmt.Printf("The square has area:%f\n", areaIntf.Area())
}
