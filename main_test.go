package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		// TODO: Add test cases.
		{"Main", 6},
		{"Solve", 0},
		{"Solve", 1},
		{"Solve", 2},
		{"Solve", 3},
		{"Solve", 4},
		{"Solve", 5},
		{"Space", 0},
		{"Space", 1},
		{"Space", 2},
		{"Space", 3},
		{"Space", 4},
		{"Space", 5},
		{"Mark", 0},
		{"Mark", 1},
		{"Mark", 2},
		{"Mark", 3},
		{"Mark", 4},
		{"Mark", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			switch tt.name {
			case "Main":
				main()
			case "Solve":
				Solve(tt.n)
				fmt.Println()
			case "Space":
				for i := 1; i <= tt.n; i++ {
					Space()
				}
				fmt.Println("<---\n")
			case "Mark":
				for i := 1; i <= tt.n; i++ {
					Mark()
				}
				fmt.Println("\n-----\n")
			default:
				fmt.Println(tt.name, "Unexpected test type called")
			}
		}) // end test func
	}
}
