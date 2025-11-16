package main

import "fmt"



func Say(animal string) string{
	if animal =="dog"{
		return "wuf"
	}else if animal =="cat"{
		return "Myau"
	} else{
		return "nah"
	}
}

func Voice(animal string, how func(string) string) {
	fmt.Println(how(animal))
}

func main() {
	Voice("dog", Say)
}

///////////lessonPractics://////////////



// var funcVar func(int) int

// func Sq(x int) int{
// 	return x*x
// }

// func main() {
// 	funcVar = Sq
// }

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