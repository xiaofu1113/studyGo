package main

import (
	"fmt"
)

func main() {
	fmt.Println("我来了，golang")

	/*var i interface{}
	i = 4
	fmt.Println(i, reflect.TypeOf(i))
	i = "abc"
	fmt.Println(i, reflect.TypeOf(i))*/

	var arr [3]*string
	arr1 := [3]*string{new(string), new(string), new(string)}
	*arr1[0] = "red0"
	*arr1[1] = "red1"
	*arr1[2] = "red2"
	arr = arr1
	fmt.Println(arr)
	var temp people
	fmt.Println(temp.getName())

}

type people struct {
	name string
	age  int
	aAge *int
}

func (peo *people) getName() string {
	return peo.name
}
