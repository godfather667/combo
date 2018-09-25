// combo.go - This a demonstration of a method to get all combinations of the elements in a slice.
// It is unique because of the method of using a binary counters to generate the combinations.
// It includes functions for both one and two dimensional arrays, which shows how to extend
// the combo function for 'n' dimensional levels.
package main

import (
	"fmt"
)

// Channel Data Structure for one dimensional slice data.
type c1data struct {
	cm   []int
	done bool
}

type c2data struct {
	cm   [][]int
	done bool
}

// main - exercises the combo function.
func main() {

	fmt.Println("Running one dimensional Test")
	c1 := run1D()
	for _, v := range c1 {
		fmt.Println(v.cm)
	}
	fmt.Println("\nTest Complete!")

	fmt.Println("Running two dimensional Test")
	c2 := run2D()
	for _, v := range c2 {
		fmt.Println(v.cm)
	}
	fmt.Println("\nTest Complete!")
}

// Run 1D test sequnce
func run1D() []c1data {
	elm := []int{100, 111, 223, 44, 552}
	cd := make(chan c1data)
	var cc []c1data
	go combo1D(cd, elm)
	for {
		c := <-cd
		if c.done == true { // Check got end of stream!
			break
		}
		cc = append(cc, c)
	}
	return cc
}

// Run 2D Test Sequence
func run2D() []c2data {
	elm := [][]int{
		{0, 5, 9},
		{1, 1, 5},
		{6, 2, 3},
		{4, 2, 4},
	}
	cd := make(chan c2data)
	var cc []c2data
	go combo2D(cd, elm)
	for {
		c := <-cd           // Channel provides stream of combinations.
		if c.done == true { // Check got end of stream!
			break
		}
		cc = append(cc, c)
	}
	return cc
}

// Combo1D - a "Combination Generator" that given a 1D slice of values
// will return via specified channel a stream of values that
// constitute all the possible combinations of the supplied slice
// of values.
func combo1D(cd chan c1data, elm []int) {
	n := len(elm)
	var lcd c1data
	lcd.done = false
	for num := 0; num < (1 << uint(n)); num++ { // The Binary Generator.
		com := []int{}
		for ndx := 0; ndx < n; ndx++ {
			if num&(1<<uint(ndx)) != 0 { // Applies Binary Mask.
				com = append(com, elm[ndx])
			}
		}
		if len(com) != 0 { // Skip Zero Result
			lcd.cm = com
			cd <- lcd
		}
	}
	lcd.done = true // Indicate end of stream!
	cd <- lcd
	return
}

// Combo2d - a "Combination Generator" that given a 2D slice of values
// will return via specified channel a stream of values that
// constitute all the possible combinations of the supplied slice
// of values.
func combo2D(cd chan c2data, elm [][]int) {
	n := len(elm)
	var lcd c2data
	lcd.done = false
	for num := 0; num < (1 << uint(n)); num++ { // The Binary Generator.
		com := [][]int{}
		lcd.cm = [][]int{}
		for ndx := 0; ndx < n; ndx++ {
			if num&(1<<uint(ndx)) != 0 { // Applys Binary Mask.
				com = append(com, elm[ndx])
			}
		}
		if len(com) != 0 { // Skip Zero Result
			lcd.cm = com
			cd <- lcd
		}
	}
	lcd.done = true // Indicate end of stream!
	cd <- lcd
	return
}
