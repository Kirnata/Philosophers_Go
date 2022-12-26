package main

import "time"

type Info struct {
	NumOfPhilos int
	TimeToDie   time.Duration
	TimeToEat   time.Duration
	TimeToSleep time.Duration
	NumOfMeals  int
}

type Fork chan struct{}

type Philo struct {
	Id         int
	MealCount  int
	RightFork  Fork
	LeftFork   Fork
	LastEating time.Time
	Info       *Info
}
