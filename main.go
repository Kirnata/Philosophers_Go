package main

import (
	"flag"
	"fmt"
	"time"
)

func NewInfo() *Info {
	return &Info{}
}

func Pukinit(inf *Info) {
	flag.IntVar(&inf.NumOfPhilos, "philosophers", 10, "The integer param")
	flag.DurationVar(&inf.TimeToDie, "die", time.Millisecond*200, "The duration param")
	flag.DurationVar(&inf.TimeToEat, "eat", time.Millisecond*50, "The duration param")
	flag.DurationVar(&inf.TimeToSleep, "sleep", time.Millisecond*50, "The duration param")
	flag.IntVar(&inf.NumOfMeals, "meals", 10, "The integer param")
	flag.Parse()
}

func main() {
	myInfo := NewInfo()
	Pukinit(myInfo)
	fmt.Println(myInfo)
	philosophers(myInfo)
}
