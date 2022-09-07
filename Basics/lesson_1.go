package main
import "fmt"

//package level every statement begins with a keyword 
//(var, func, and so on)

func main() {
	// 1.
	// fmt.Println(add(1,2))
	
	// 2.
	// resultOne,resultTwo := swap("Hello","World") // short declare
	// fmt.Println(resultOne,resultTwo)

	// 3.
	// x,y := naked(1)
	// fmt.Println(x,y)

	// 4. 
	// var (
	// 	boolVar bool = true
	// 	intVar int = 3
	// 	stringVar string = "hello"
	// )
	// fmt.Printf("%v,%v,%v",boolVar,intVar,stringVar)

	// 5. 
	// var a int
	// var b float32 
	// var c bool 
	// var d string 
	// fmt.Printf("%v,%v,%v,%q\n",a,b,c,d)
	
	// 6.
	// var a int = 10
	// fmt.Printf("type: %T\n",float32(a))

	// 7.
	// const a = "3"//Constants cannot be declared using the := syntax.
	// fmt.Println("裡面有",a,"種水果")
	// fmt.Printf("裡面有%v種水果\n",a)

	// 8.
	var a int = 1
	fmt.Println(a<<1)
	fmt.Println(a<<62)
}

// function
func add(a /*int*/ ,b int) int {
	return a+b
}
// multiple results
func swap(a ,b string) (string,string) {
	return b,a
}

// naked return
func naked(base int) (x,y int) {
	x = base + 2
	y = base + 3
	return
}

/*
- Go 屬於強型別（Static Types）的語言
ref: https://pjchender.dev/golang/variables/

	- variables
		x int
		p *int
		a [3]int
		
		---
		1. var a int 
			a = 3

		2. a := 3

		3. var a,b,c int

		4. var a,b,c int = 1,2,3

		5. var a,b,c = 3,true,"apple" -> If an initializer is present, the type can be omitted;
	- types

		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // alias for uint8

		rune // alias for int32
			// represents a Unicode code point

		float32 float64

		complex64 complex128


*/

/* 
- function
 func
*/