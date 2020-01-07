package main

import "fmt"

func add(x int, y int)int {
	return x+y
}

func swap(x int, y string) (string,int) {
	return y,x
}

func named_return(x int, y string) (a int, b string){
	a = x+1
	b = y + "lols"
	return
}

var a int
var b int = 1
var c = 2

func main() {
	// fmt.Println(add(1,20))
	// fmt.Println(swap(1,"20"))
	// fmt.Println(named_return(1,"haha"))
	//cannot be outside fn
	// d := 40
	// fmt.Println(d)
	// e := float64(d)
	// fmt.Println(e+0.01)
	// f := string(d)
	// fmt.Println(f)
	const pi = 3.14
	fmt.Println(pi)
}
