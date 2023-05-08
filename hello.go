package main

import (
	"fmt"
	"example.com/hello/src"
)

var global_num int = 69

func main() {
    var num = 42
	course1 := "CSE 312"
	
	fmt.Println(num)
	fmt.Println(course1, "is the string")
	
	print_global()
	
	fmt.Println(src.Add(10, 20))
	
	if num == 42 {
		fmt.Printf("The string is %s\n", course1)
	}
	
	fmt.Printf("%b\n", num)
	fmt.Println(5 > 4)
	
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	
	numsArr := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numsArr)
	for index, val := range numsArr {
		fmt.Println(index, val)
	}
	fmt.Println("Thje length of the array is", len(numsArr))
}

func print_global() {
	fmt.Println(global_num)
}
