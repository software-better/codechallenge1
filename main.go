package main

import "fmt"

func main() {
	// At first I thought I would assemble a matrix…
	// But it turned out that the problem could be
	// solved procedurally, and the matrix approach seemed
	// trickier than expected
	//
	// Inputs
	// 		n = number of rows & columns
	// Output
	// 		n rows, right justified with marks preceded by
	//		spaces
	// Process
	// 		for each row, starting at the top, print the
	// 		required number of spaces and marks followed
	//		by a newline

	n := 6 // matrix dimensions

	// At first I thought I would assemble a matrix…
	matrix := make([]string, 10)
	matrix[0] = "Hello"
	matrix[1] = "World"
	fmt.Println(matrix[0], matrix[1])

	// i represents which row we are working on,
	// starting at the top
	//
	// n is “input” specified for the algorithm (how many
	// rows will exhibit this behavior
	//
	// marks and spaces are our target numbers for each row
	// spaces are always printed first, followed by the marks

	Solve(n)
}

func Solve(n int) {
	for i := 1; i <= n; i++ {
		// s := make([]string, n)
		spaces := n - i
		marks := n - spaces
		for j := 0; j <= (spaces - 1); j++ {
			Space()
		}
		for j := 0; j <= (marks - 1); j++ {
			Mark()
		}
		fmt.Println()
	}
}

func Space() {
	fmt.Print(" ")
}
func Mark() {
	// s = append(s, "X")
	fmt.Print("X")
}

/*******************
[ `go run main.go` | done: 234.368687ms ]
	Hello World
	      X
	     XX
	    XXX
	   XXXX
	  XXXXX
	 XXXXXX
[ ~/Go/src/codechallenge/ ] #
*******************/
