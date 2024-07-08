package main

import "fmt"

func returnValue() {

	var caugthValue string = returnsValue("return value from function")

	fmt.Println(caugthValue)

	someString, someInt, someBool := returnsMultiple()

	fmt.Println(someString, someInt, someBool)

}

func returnsValue(somethig string) string {
	return somethig
}

func returnsMultiple() (string, int, bool) {
	return "string", 1, true
}
