package main

import "fmt"

const Loginname = "Mac"

func main() {

	var Username string = "Macruz"
	fmt.Println(Username)
	fmt.Printf("variable type: %T \n", Username)

	var Id uint8 = 00255
	fmt.Println(Id)
	fmt.Printf("variable type: %T \n", Id)

	var logged bool = true
	fmt.Println(logged)
	fmt.Printf("variable type: %T \n", logged)

	var website = 567
	fmt.Println(website)
	fmt.Printf("Variable type: %T \n", website)

	Quantity := 5
	fmt.Println(Quantity)
	fmt.Printf("Variable type: %T \n", Quantity)

	fmt.Println(Loginname)
	fmt.Printf("Variable type: %T \n", Loginname)

}
