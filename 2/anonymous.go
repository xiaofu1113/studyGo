package main

import (
	"fmt"
	"time"
)

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

func main() {
	fmt.Println(time.Now())
	go func(str string) {
		fmt.Println("看谁快!!", str)
	}("我的")
	people1 := people{}
	fmt.Println(people1.action("人很牛逼了"))
	time.Sleep(1 * time.Second)
	people2 := yellowRace{}
	fmt.Println(people2.action("黄种人也牛逼了"))
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now())
}
