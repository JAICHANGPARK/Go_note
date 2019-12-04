package main

import "fmt"

func main() {

	rawLiteral := `아리랑\n아리랑\n아라리요`
	interLiteral := "아리랑아리랑\n아리리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)

}
