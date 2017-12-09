// https://gobyexample.com/
//https://learnxinyminutes.com/docs/go/ 

package main

import "fmt"

func add(a int, b int) int {
	x := 10
	return x + a
}

func main() {
		fmt.Println("Hello World")
		// x := 5
		// y := 5
		// sum, prod := learnMultiple(x, y)
		x := []int{1,2,3,4,5}
		x = append(x, []int{1,2,3,4,5}...)
		// literal := `
		// 	wow this is a literal!

		// 	yes
		// `
		// fmt.Println(x)

		// p, q := learnMemory() // Declares p, q to be type pointer to int.
		// fmt.Println(*p, *q)   // * follows a pointer. This prints two ints.

		z := learnNamedReturns(3, 4)

		fmt.Println(z)
		

}

func learnMultiple(x, y int) (sum, prod int){
	return x + y, x * y
}


func learnMemory() (p, q *int) {
	// Named return values p and q have type pointer to int.
	p = new(int) // Built-in function new allocates memory.
	// The allocated int is initialized to 0, p is no longer nil.
	s := make([]int, 20) // Allocate 20 ints as a single block of memory.
	fmt.Println(s)
	s[3] = 7             // Assign one of them.
	fmt.Println(s)
	r := -2              // Declare another local variable.
	return &s[3], &r     // & takes the address of an object.
}

func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return // z is implicit here, because we named it earlier.
}
