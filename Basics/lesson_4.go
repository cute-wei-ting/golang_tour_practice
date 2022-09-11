/* Methods and interfaces */
package main 
import (
	"fmt"
	"math"
	// "strconv"
	// "golang.org/x/tour/reader"
	"io"
	// "os"
	// "strings"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Vertex struct {
	X,Y float64
}
// a method is just a function with a receiver argument.
func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func AbsFn(v Vertex) float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func (v *Vertex) Scale(s float64) {
	v.X = v.X * s
	v.Y = v.Y * s
}

func ScaleFn(v *Vertex,s float64)  {
	v.X = v.X * s
	v.Y = v.Y * s
}

type CustomFloat float64
func (f CustomFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
//
type Abser interface {
	Abs2() float64
}
func (v *Vertex) Abs2() float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}
func (f CustomFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
//
type I interface {
	M()
	N()
}
type T struct {
	X string
}
func (t T) M() {
	fmt.Println(t.X)
}
func (t *T) N() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.X)
}
func describe(i I) {
	fmt.Printf("(%v,%T)\n",i,i)
}
func describe2(i interface{}) {
	fmt.Printf("(%v,%T)\n",i,i)
}
//
func printType(i interface{}) {
	switch c:=i.(type) {
		case int:
			describe2(c)
		case float64:
			describe2(c)
		default:
			fmt.Printf("I don't know about type %T!%v\n", c,c)

	}
}
//
type Person struct {
	Name string
	Age int
}
func (p Person) String() string {
	return fmt.Sprintln("Name:",p.Name,",Age:",p.Age)
}
//
type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	result:=""
	for i,v:=range ipaddr {
		if i != 0 {
			result+=","
		}
		result+=fmt.Sprint(v)
	}
	return result
}
// 
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintln("cannot Sqrt negative number:",float64(e))
}

func Sqrt (x float64) (float64, error) {
	if x < 0 {
		return 0,ErrNegativeSqrt(x)
	}
	return math.Sqrt(x),nil
}
//
type MyReader struct{}

func (r MyReader) Read(b []byte) (int,error) {
	for i:=range b {
		b[i]='A'
	}
	return len(b),nil
}
//
type rot13Reader struct {
	r io.Reader
}
func (rot *rot13Reader) Read(b []byte) (int,error) {
	n,e:=rot.r.Read(b)
	for i:=0;i<n;i++ {
		if (b[i] >= 'A' && b[i] <= 'N') || (b[i] >= 'a' && b[i] <= 'n') {
			b[i]+=13
		}
		if (b[i] >= 'M' && b[i] <= 'Z') || (b[i] >= 'm' && b[i] <= 'z') {
			b[i]-=13
		}
	}
	return n,e
}
//
type Image struct {
	w, h int
	c    uint8
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) At(x, y int) color.Color {
	return color.RGBA{img.c + uint8(x*y+x*y), img.c + uint8(x*y+x*y), 255, 255}
}
func main() {
	// 1. Method
		// a:=Vertex{3,4}
		// fmt.Println(a.Abs())
		// b:=Vertex{6,8}
		// fmt.Println(b.Abs())
	// 2. 
		// a:=CustomFloat(-math.Sqrt2)
		// fmt.Println(a)
		// b:=CustomFloat(10)
		// fmt.Println(b)
	// 3. Pointer receivers
		/* With a value receiver, the Scale method operates on a copy of the original Vertex value.
		 (This is the same behavior as for any other function argument.) 
		 The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
		*/
		// a:=Vertex{3,4}
		// a.Scale(10)
		// fmt.Println(a.Abs())
	// 4-1. Methods and pointer indirection
		/* interprets the statement a.Scale(5) as (&a).Scale(5) */

		// a:=Vertex{3,4}
		// a.Scale(10)
		// ScaleFn(&a,2)
		// fmt.Println(a)

		// b:=Vertex{6,8}
		// p:=&b
		// p.Scale(10)
		// ScaleFn(p,2)
		// fmt.Println(b)
	// 4-2. Methods and pointer indirection(2)
		// The equivalent thing happens in the reverse direction.
		/* b.Abs() is interpreted as (*b).Abs(). */
		// a:=Vertex{3,4}
		// fmt.Println(a.Abs())
		// fmt.Println(AbsFn(a))

		// b:=&Vertex{6,8}
		// fmt.Println(b.Abs())
		// fmt.Println(AbsFn(*b))
	// 5-1. Interfaces
		// var a Abser
		// b:=Vertex{3,4}
		// c:=CustomFloat(1.414)
		// a=c
		// fmt.Println(a.Abs2()) 
		/* error
		   a=b
		   Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
		*/
		// a=&b
		// fmt.Println(a.Abs2()) 
	// 5-2.
		// var a I = T{"I'm X"}
		// a.M()
	// 5-3. Interface values with nil underlying values
		// var a I = &T{"I'm Here"}
		// a.M()
		// a.N()

		// var t *T
		// var b I = t
		// /* exception
		//    b.M()
		// */
		// b.N()
	// 5-4. Nil interface values 
		// var i I
		// describe(i)
		/* exception 
			Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
			i.M()
		*/
	// 6. The empty interface
		// var emp interface{}
		// describe2(emp)

		// emp = 3
		// describe2(emp)

		// emp = "HELLO"
		// describe2(emp)
	// 6-1. Type assertions

		/* To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
			t, ok := i.(T)
		*/
		// var i interface{}
		// i="HAPPY"
		// s,ok:=i.(string)
		// fmt.Println(s,ok)

		// f,ok:=i.(int)
		// fmt.Println(f,ok)

		/* panic 
			f=i.(int)
			fmt.Println(f)
		*/
	// 6-2. Type switches
		// var i interface{}
		// i=3
		// printType(i)
		// i="hehe"
		// printType(i)
		// i=-1.33
		// printType(i)
	// 7-1. Stringers
		// p1:=Person{"John",20}
		// p2:=Person{"Mary",30}
		// fmt.Println(p1,p2)
	// 7-2. Exercise: Stringers
		// hosts:=map[string]IPAddr{
		// 	"personal": {127,0,0,1},
		// 	"company": {8,8,8,8},
		// }
		// for name,ip:=range hosts {
		// 	fmt.Println(name,ip)
		// }
	// 8-1. Errors
		// i,error:=strconv.Atoi("42")
		// if error != nil {
		// 	fmt.Printf("Couldn't convert %v\n",error)
		// 	return
		// }
		// fmt.Printf("Convert Success %v\n",i)
	// 8-2. Errors
		// fmt.Println(Sqrt(2))
		// fmt.Println(Sqrt(-2))
	// 9. Exercise: Readers
		// reader.Validate(MyReader{})
	// 10. Exercise: rot13Reader
		// s:=strings.NewReader("Lbh penpxrq gur pbqr!")
		// r:=rot13Reader{s}
		// io.Copy(os.Stdout,&r)
	// 11. Exercise: Images
		m := Image{100, 100, 128}
		pic.ShowImage(&m)
}

