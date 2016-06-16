package main

import "fmt"

type  inners struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	inners
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is :%d\n", outer.b)
	fmt.Printf("outer.c is :%d\n", outer.c)
	fmt.Printf("outer.int is :%d\n", outer.int)
	fmt.Printf("outer.in1 is :%d\n", outer.in1)
	fmt.Printf("outer.in2 is :%d\n", outer.in2)

	outer2 := outerS{6, 7.5, 60, inners{5, 10}}
	fmt.Printf("outer2 is :", outer2)
}