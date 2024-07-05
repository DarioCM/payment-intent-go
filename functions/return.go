package main

import "fmt"

func returnValue() {

	var caugthValue string = returnsValue("return value from function")

	fmt.Println(caugthValue)

}

func returnsValue(somethig string) string {
	return somethig
}
