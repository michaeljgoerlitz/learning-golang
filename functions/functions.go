package functions

import "fmt"

//	func sub(x int, y int) int {
//		return x - y
//	}
func HelloWorld() string {
	return "hello world"
}

func Subtract(x int, y int) int {
	return x - y
}

func Addition(x, y int, template string) string {
	sum := x + y
	return template + fmt.Sprintf("%d", sum)
}

func IncrementNoReturn(x int) {
	x++
}

func IncrementReturn(x int) int {
	x++
	return x
}

func TestIncrements() {
	x := 5
	IncrementNoReturn(x)
	fmt.Println(x) // should still be 5

	x = IncrementReturn(x)
	fmt.Println(x) // should be 6

}
