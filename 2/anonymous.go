package main

import "fmt"

type people struct {
	name    string
	age     int
	address string
}

func (people *people) action(acrion string) string {

	return acrion
}

type yellowRace struct {
	skin string
	people
}

func main()  {
	people1 :=people{}
	fmt.Println( people1.action("人很牛逼了"))
	people2 :=yellowRace{}
	fmt.Println(people2.action("黄种人也牛逼了"))
}