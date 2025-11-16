package main

import "fmt"




func main() {
	foo()
}

///////////lessonPractics://////////////

func foo(){
	defer func(){
		r := recover()
		if r != nil{
			fmt.Println("поймана паника")
		}
	}()
	panic("это паника")
}

func def(){
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}
var name string

func init(){
	name = "Evgeniy"
}

func init(){
	name ="Arsenya"
}

// зарезервированные имена:
// нет сигнатуры, не библиотека, встроенные в язык функции
// make(), new(), len(), cap(), delete(), close(), append(), copy()
// panic(), recover(), main(), 
// init() - вызывается один раз при инициализации пакета автоматически, до функции main,


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

// func main() {
// 	Voice("dog", Say)
// }


// var funcVar func(int) int

// func Sq(x int) int{
// 	return x*x
// }

// func main() {
// 	funcVar = Sq
// }

func Input2ints (x int, y int){
	fmt.Printf("Два числа: %d и %d", x, y )
}

func Returns2ints()(x int, y int){
	return 7, 10;
}

func TheirdFunc(){
	Input2ints(Returns2ints())
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