/* More types: structs, slices, and maps. */

package main
import (
	"fmt"
	// "golang.org/x/tour/pic"
	// "golang.org/x/tour/wc"
	"strings"
	// "math"
)

type Vertex struct {
	X int 
	Y int
}

func main() {
	// 1. pointer
		// i,j:=10,30

		// iPointer:=&i
		// fmt.Println(*iPointer)
		// *iPointer = 101
		// fmt.Println(i)

		// jPointer:=&j
		// fmt.Println(*jPointer)
		// *jPointer = *jPointer /10
		// fmt.Println(j)
	// 2-1. struct
		// v:=Vertex{3,4}
		// fmt.Println(v.Y)

		// p:=&v 

		//To access the field X of a struct when we have the struct pointer p we could write (*p).X. 
		//However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
		
		// p.X=6
		// fmt.Println(v.X)
	// 2-2.
		// v:=Vertex{X:8,Y:9}
		// v1:=Vertex{X:7}
		// v2:=Vertex{Y:9}
		// p:=&Vertex{}
		// fmt.Println(v)
		// fmt.Println(v1)
		// fmt.Println(v2)
		// fmt.Println(p)
		// fmt.Println(*p)
	// 3-1. arrays
		// var a [2]string
		// a[0] = "1"
		// a[1] = "2"
		// fmt.Println(a)
		// fmt.Println(a[1])

		// b:= [5] int {5,4,3,2,1} 
		// fmt.Println(b)
		// fmt.Println(b[1])
	// 3-2.	Slices are like references to arrays
		// arr := [5]int{1,2,3,4,5}
		// a := arr[0:3]
		// b := arr[1:4]
		// fmt.Println(a)
		// fmt.Println(b)

		// b[0] = 10
		// fmt.Println(arr) 
	// 3-3. A slice literal is like an array literal without the length.
		// This is an array literal: [3]bool{true, true, false}
		// And this creates the same array as above, then builds a slice that references it: []bool{true, true, false}

		// a := []string{ "hello", "I'am","skye"}
		// fmt.Println(a)

		// b := []struct{
		// 	X int
		// 	Y int
		// }{
		// 	{1,2},
		// 	{3,4},
		// }
		// fmt.Println(b)
		// fmt.Println(b[0])
	// 3-4. Slice length and capacity , 
		// The length of a slice is the number of elements it contains.
		// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. https://go.dev/tour/moretypes/11
		// You can extend a slice's length by re-slicing it, provided it has sufficient capacity
			// example := []int {8,7,6,5,4}
			// example = example[:0]
			// printSlice(example)
			// example = example[2:4] // drop its first two values.
			// printSlice(example)
			// example = example[0:3]
			// printSlice(example)
	// 3-5. Nil slices
		// The zero value of a slice is nil.
		// A nil slice has a length and capacity of 0 and has no underlying array.
		// var a []int
		// printSlice(a)
		// if(a == nil){
		// 	fmt.Println("this is nil array")
		// }
	// 3-6. append
		// var a []int 
		// printSlice(a)
		// a=append(a,2)
		// printSlice(a)
		// a=append(a,3,4,5)
		// printSlice(a)
		// a=append(a,6)
		// printSlice(a)
	// 3-7. range 
		// a:=[]int { 12,24,36,48,60,72 }
		// for i,v:=range a {
		// 	fmt.Printf("index: %d, value: %v\n",i,v)
		//}
	// 3-8. make , range
		//  arr:=make([]int,10)
		// for i:=range arr {
		// 	arr[i] = 1 << i
		// }
		// for _,v:=range arr {
		// 	fmt.Println(v)
		// }
	// 3-9. Exercise: Slices
		//pic.Show(Pic)
	// 4-1. Maps
		// var a map[string]Vertex
		// a = make(map[string]Vertex)//init
		// fmt.Println(a)
		// a["skye"] = Vertex{2,3}
		// fmt.Println(a["skye"])
	// 4-2.
		// a:=map[string]Vertex{
		// 	"HELLO" : Vertex{1,2},
		// 	"WORLD" : Vertex{3,4},
		// }
		// fmt.Println(a["HELLO"])
		// fmt.Println(a["WORLD"])
		// a:=map[string]Vertex{
		// 	"HELLO" : {1,2},
		// 	"WORLD" : {3,4},
		// }
		// fmt.Println(a["HELLO"])
		// fmt.Println(a["WORLD"])
	// 4-3. special use
		// a:=make(map[string]int)
		// a["John"] = 23
		// a["Tim"] = 30
		// fmt.Println(a)

		// delete(a,"Tim")
		// fmt.Println(a)

		// elem,ok:=a["John"]
		// fmt.Println(elem,ok)

		// elem,ok=a["Tim"]
		// fmt.Println(elem,ok)
	// 5.
		// wc.Test(WordCount)
	// 6. pass function as variable
		// fmt.Println(compute(math.Pow,3,4))
		// temp:=func(x,y float64) float64{
		// 	return math.Sqrt(x*x+y*y)
		// }
		// fmt.Println(compute(temp,3,4))
		// fmt.Println(temp(3,4))
	// 7. closure
		// pos,neg := adder(),adder()
		// for i:=0;i<10;i++ {
		// 	fmt.Println(pos(i),neg(-i))
		// }
	// 8. Exercise: Fibonacci closure
		// fn:=fibonacci()
		// for i:=0;i<10;i++ {
		// 	fmt.Println(fn())
		// }
}

func printSlice(a []int) {
	fmt.Printf("len: %d , cap: %d, array: %v\n",len(a),cap(a),a)
}

func Pic(dx, dy int) [][]uint8 {
	// Allocate two-dimensioanl array.
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}
	
	// Do something.
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch {
			case j % 15 == 0:
				a[i][j] = 240
			case j % 3 == 0:
				a[i][j] = 120
			case j % 5 == 0:
				a[i][j] = 150
			default:
				a[i][j] = 100
			}
		}
	}	
	return a
}
func WordCount(s string) map[string]int {
	arr:=strings.Fields(s)
	m:=make(map[string]int)
	
	for _,v:=range arr {
		m[v]++
	}
	return m
}

func compute(fn func(float64,float64) float64,a float64 ,b float64) float64{
	return fn(a,b)
}

func adder() func(int) int {
	sum:=0
	return func(x int) int {
		sum+=x
		return sum
	}
}

func fibonacci() func() int {
	count:=-1
	before:=0
	now:=1

	return func() int {
		count++
		if count == 0 || count == 1 {
			return count
		}
		sum:=before + now
		before = now
		now = sum
		return sum
	}
}