/* Flow control statements: for,if,else,switch and defer */

package main

import (
	"fmt"
	"math"
)

func main() {
	// 1.
	//  for i:=0;i<10;i++ {
	// 	fmt.Println(i)
	//  }

 	// 2.
	// i:=2 
	// for ;i<10;i++ {
	// 	fmt.Println(i)
	// }

	// 3.  while in GO
	//  i:=5
	//  for i<10 {
	//  	fmt.Println(i)
	// 	i+=1
	//  }

	// 4. about print function
	// name:="apple"
	// age:=3
	// fmt.Print(name,"is",age,"years old.\n")
	// fmt.Printf("%v is %v years old.\n",name,age)
	// fmt.Println(name,"is",age,"years old.")

	// 5. sqrt function practice if statement
	// fmt.Println(sqrt(10))
	// fmt.Println(sqrt(4))
	// fmt.Println(sqrt(-4))
	
	// 6-1. [if] with a short statement, and powOrLimit function
		// fmt.Println(powOrLimit(3,3,28)) //27
		// fmt.Println(powOrLimit(3,3,6)) //6
		// fmt.Println(powOrLimit(2,4,17)) //16
		// fmt.Println(powOrLimit(2,4,15)) //15
	// 6-2. Variables declared inside an if short statement are also available inside any of the else blocks.
		// fmt.Println(powWithIfElse(3,3,28))
		// fmt.Println(powWithIfElse(3,3,6))
		// fmt.Println(powWithIfElse(2,4,17)) 
		// fmt.Println(powWithIfElse(2,4,15)) 
	// 7-1. switch
		// var num int
		// fmt.Println("please input a number")
		// fmt.Scanln(&num)

		// switch num {
		// 	case 1:
		// 		fmt.Println("this is one")
		// 	case 2:
		// 		fmt.Println("this is two")
		// 	default:
		// 		fmt.Println("this is default")
		// }
		//7-2. switch without a condition ,This construct can be a clean way to write long if-then-else chains.
		// var num int
		// fmt.Println("please input a number")
		// fmt.Scanln(&num)

		// switch {
		// 	case num==1:
		// 		fmt.Println("this is one")
		// 	case num==2:
		// 		fmt.Println("this is two")
		// 	default:
		// 		fmt.Println("this is default")
		// }
	// 8-1. defer
		// defer fmt.Println("world")

		// fmt.Println("hello")
		// fmt.Println("befor")
	// 8-2. Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
		// fmt.Println("begin")
		// for i:=1;i<=10;i++ {
		// 	defer fmt.Println(i)
		// }
		// fmt.Println("done")
}

//
func sqrt (num float64) string {
	if num < 0 {
		return sqrt(-num) + "i"
	}
	return fmt.Sprint(math.Sqrt(num))
} 

func powOrLimit (x,n,lim float64) float64 {
	if v:=math.Pow(x,n);v<lim {
		return v
	} 
	return lim
}

func powWithIfElse (x,n,lim float64) float64 {
	if v:=math.Pow(x,n);v<lim {
		fmt.Println(v,"<",lim)
		return v
	} else {
		fmt.Println(v,">=",lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := x/2
	for i:=0;i<10;i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2*z)
	}
	return z
}

