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

	for ; i < 5; i++ {
		println(i)
	}

	for j := 0; j < 3; j++ {
		println(j)
	}

	// Setting up infinite loop
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}
}
