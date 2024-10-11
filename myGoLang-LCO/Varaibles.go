package main

import "fmt"

const LoginToken string = "abcdefgh"

func main() {
	var username string = "Mohan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T\n", isLoggedin)
	/// uint8 = 0 to 255
	var sval uint8 = 255
	fmt.Println(sval)
	fmt.Printf("Variable is of type: %T\n", sval)
	//float
	var svalFloat float32 = 255.2551234
	fmt.Println(svalFloat)
	fmt.Printf("Variable is of type: %T\n", svalFloat)

	// default values and some aliases
	// no garbage value sif it is not intialized
	var anotherVariables int
	fmt.Println(anotherVariables)
	fmt.Printf("Variable is of type: %T\n", anotherVariables)

	//implict type
	var website = "https://www.google.com"
	fmt.Println(website)
	//  we we dosn;t give th e variables type,then it will take by default
	// and we if want  to change to other type it will not allow it.
	//var website = 23
	//fmt.Println(website)

	// no var style
	numsOfUser := 300000.0 // in side a method Only we can  Use.
	fmt.Println(numsOfUser)
	fmt.Println("\n")
	////print the public varible
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)

}
