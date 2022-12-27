package main

import (
	"flag"
	"time"
)

func newInfo() *Info {
	return &Info{}
}

func (info *Info) init() {
	flag.IntVar(&info.NumOfPhilos, "philosophers", 10, "The integer param")
	flag.DurationVar(&info.TimeToDie, "die", time.Millisecond*200, "The duration param")
	flag.DurationVar(&info.TimeToEat, "eat", time.Millisecond*50, "The duration param")
	flag.DurationVar(&info.TimeToSleep, "sleep", time.Millisecond*50, "The duration param")
	flag.IntVar(&info.NumOfMeals, "meals", 10, "The integer param")
	flag.Parse()
}

func main() {
	myInfo := newInfo()
	myInfo.init()
	//fmt.Println(myInfo)
	//return
	philosophers(myInfo)
}
