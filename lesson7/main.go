package main

import "fmt"

func Cube(x int) (res int) {
	res = x*x*x
	return res
}

func main() {
	res := Cube(2)
	fmt.Println(res)
}