package main

import (
	"fmt"
	"os"
	"time"
	//"strconv"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is:%s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	t := time.Now()
	fmt.Println(t)

	var i1 = 5
	fmt.Printf("An integer:%d,it's location in memory:%p\n", i1, &i1)

	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	//var orgi string = "ABC"
	//
	//var newS string
	//fmt.Printf("The size of ints is : %d\n", strconv.IntSize)
	//an, err := strconv.Atoi(orgi)
	//if err != nil {
	//	fmt.Printf("orgi %s is not an integet - exiting with error\n", orgi)
	//	return
	//}
	//fmt.Printf("The integer is %d\n", an)
	//an = an + 5
	//newS = strconv.Itoa(an)
	//fmt.Printf("The new string is :%s\n", newS)

	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d interation \n", i)
	}

	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

	LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	fmt.Println("A statement just after for loop.")


}


