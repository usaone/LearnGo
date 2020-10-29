package main

import (
	"fmt"
)

func collections() {
	arrays()
	slices()
	slices2()
	maps()
	structs()
}

// begin arrays
func arrays() {
	fmt.Println(">>>>> Begin arrays <<<<<")
	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{3, 4, 5}
	fmt.Println(arr2)
	fmt.Println("Length of array: ", len(arr2))
}

// end arrays

// begin slices
func slices() {
	fmt.Println(">>>>> Begin slices <<<<<")
	arr := [3]int{1, 3, 5}

	slice := arr[:]

	fmt.Println(arr, slice)

	arr2 := [5]int{4, 5, 6, 7, 8}
	slice2 := arr2[1:4] // returns [5 6 7]
	fmt.Println(arr2, slice2)

	slice2[0] = 40
	fmt.Println(arr2, slice2)
	fmt.Println("Length of slice: ", len(slice2))
}

// end slices

// begin slices2
func slices2() {
	fmt.Println(">>>>> Begin slices2 <<<<<")
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4, 9, 11, 17)
	fmt.Println(slice)
	fmt.Println("Length of slice: ", len(slice))
}

// end slices2

// begin maps
func maps() {
	fmt.Println(">>>>> Begin maps <<<<<")
	m := map[string]int{"foo": 42, "coo": 94}
	fmt.Println(m)
	fmt.Println(m["coo"])
	m["coo"] = 29
	fmt.Println(m)

	j := map[int]string{3: "axis", 5: "Bike", 2: "age", 1: "apple", -1: "jackass"}
	fmt.Println(j)
	delete(j, -1)
	fmt.Println(j)
	fmt.Println("Length of map: ", len(j))
}

// end maps

// begin structs
func structs() {
	fmt.Println(">>>>> Begin structs <<<<<")
}

// end structs
