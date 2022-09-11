/* Generics */
package main
import "fmt"

func Index[T comparable](arr []T,find T) int {
	for i,v:=range arr {
		if find == v {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val T
}
func main() {
	// 1.  generic functions
		// a:=[]int{1,2,3,4,5}
		// b:=[]string{"apple","banana"}
		// fmt.Println(Index(a,2))
		// fmt.Println(Index(b,"HELLO"))
	// 2.  generic types
		// var a,b List[int]
		// a.next=&b
		// a.val=3
		// b.val=4
		// fmt.Println(a.next.val)
}