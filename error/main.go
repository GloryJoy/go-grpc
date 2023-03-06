package main

import (
	"fmt"
	"log"
)

func main() {
	customeError, err := ErrorTest()

	if err != nil {
		log.Fatalf("We run into a type of error and the message is %s", err.Error())
	}

	fmt.Println(customeError)

}

func ErrorTest() (string, error) {
	e := NewCustomError()
	return "Hello", e
}
