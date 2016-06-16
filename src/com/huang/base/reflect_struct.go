package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1,s2,s3 string
}

func (n NotknownType) String() string{
	return n.s1+"-"+n.s2+"-"+n.s3
}

var secret interface{} = NotknownType{"Ada","Go","Oberon"}

func main() {

	value := reflect.ValueOf(secret)

	typ := reflect.TypeOf(secret)

	fmt.Println(typ)

}
