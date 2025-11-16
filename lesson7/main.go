package main

import "fmt"

var funcVar func(int) int

func Sq(x int) int{
	return x*x
}

func main() {
	funcVar = Sq
}

///////////lessonPractics://////////////

func input2ints (x int, y int){
	fmt.Printf("Два числа: %d и %d", x, y )
}

func returns2ints()(x int, y int){
	return 7, 10;
}

func theirdFunc(){
	input2ints(returns2ints())
}

func Sum(x ...int) (res int){
	for _, v := range x{
		res = v
	} 
	return 
}

func Cube(x int) (res int) {
	res = x*x*x
	return res
}