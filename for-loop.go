package main

func forTest() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			continue
		}
		if i == 5 {
			continue
		}
		println("continuing...")
	}
}
