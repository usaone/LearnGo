package main

import (
	"fmt"
)

// start by typing fm and intellisense will show fmmain which when selected
// will populate the func main bare structure.
func main() {
	fmt.Println(">>>>>+++++++ begin main +++++++<<<<<")

	port := 3000
	isStarted := startWebServer(port, 5)
	fmt.Println(isStarted)

	fmt.Println(">>>>>++++++++ end main ++++++++<<<<<")
}

func startWebServer(port, numberOfRetries int) bool {
	fmt.Println("Starting web server...")
	// do startup things
	fmt.Println("Server started : Listening on port", port)
	fmt.Println("Number of tries", numberOfRetries)
	return true
}
