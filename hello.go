package hello

import "fmt"

//Hello prints given string on stdout
func Hello(s string) {
	fmt.Println(s)
}

//HelloNew prints hello message
func HelloNew() {
	fmt.Println("Hello, world")
}

func newFunc() {
	fmt.Println("this func is useless as not exported an not called anywhere in this file")
}
