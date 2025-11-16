package main

import "fmt"

func Cube(x int) (res int) {
	res = x*x*x
	return res
}

func main() {
	res := Sum(1, 2, 3, 4)
	fmt.Println(res)
}


func Sum(x ...int) (res int){
	for _, v := range x{
		res = v
	} 
	return 
}