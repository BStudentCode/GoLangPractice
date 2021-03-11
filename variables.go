package main

import "fmt"

var (
	actorName    string = "Actor Name"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)
var (
	I int = 1 /*UPPERCASE var = global*/
)

func main() {
	var i int = 42 //One way to initialise, variable only visible inside this code block
	//	f := 42        //another way to initialise
	//	var ii int
	//	ii = 42
	fmt.Printf("%v, %T", i, i) //print with format, v = value T =
	var j float32 = 27
	fmt.Println(j)

	var a float32 = 42.5
	fmt.Printf("%v, %T\n", a, a) //There is documentation for Printf, the variable is passed twice because we used two % flags

	var b int
	b = int(a) //CASTING
	fmt.Printf("%v, %T\n", b, b)
}
