package main

import (
	"fmt"
)

func primitives() {
	variables()
	pointers()
	addressof()
	constants()
	iotas()
}

// begin variables
func variables() {
	fmt.Println(">>>>> Begin variables <<<<<")
	var i int
	i = 20
	fmt.Println(i)

	fmt.Println("Hello from a module, Gophers!")

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(7, 8)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}

// end variables

// begin pointers
func pointers() {
	fmt.Println(">>>>> Begin pointers <<<<<")
	var firstName1 *string = new(string)

	*firstName1 = "Arthur"

	fmt.Println(*firstName1)
}

// end pointers

// begin addressof
func addressof() {
	fmt.Println(">>>>> Begin addressof <<<<<")
	firstName2 := "Arthur"
	fmt.Println(firstName2)

	ptr := &firstName2 // addressof specified by &
	fmt.Println(ptr, *ptr)

	firstName2 = "Novell"
	fmt.Println(ptr, *ptr)
}

// end addressof

// begin constants
func constants() {
	fmt.Println(">>>>> Begin constants <<<<<")
	const pi = 3.1415
	fmt.Println(pi)

	const c int = 3
	fmt.Println(c + 3)

	fmt.Println(float32(c) + 1.2)
}

// end constants

// begin iotas
const (
	first = iota
	second
)

const (
	third = iota
	fourth
)

func iotas() {
	fmt.Println(">>>>> Begin iotas <<<<<")
	fmt.Println(first, second, third, fourth)
}

// end iotas
