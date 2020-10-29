package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

// start by typing fm and intellisense will show fmmain which when selected
// will populate the func main bare structure.
func main() {
	fmt.Println(">>>>>+++++++ begin main +++++++<<<<<")
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)

	fmt.Println(">>>>>++++++++ end main ++++++++<<<<<")
}
