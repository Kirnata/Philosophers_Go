package main

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var logger = log.New()

func createForks(info *Info) []Fork {
	forks := make([]Fork, info.NumOfPhilos)
	for i := range forks {
		forks[i] = make(Fork, 1)
	}
	return forks
}

func closeForks(forks []Fork) {
	for _, v := range forks {
		close(v)
	}
}

func (info *Info) start(forks []Fork) {
	var wg sync.WaitGroup
	wg.Add(info.NumOfPhilos)
	logger.Info("Starting ...")

	ctx, cancel := context.WithCancel(context.Background())
	philos := make([]*Philo, info.NumOfPhilos)
	for i := range philos {
		philos[i] = &Philo{
			Id:         i + 1,
			MealCount:  info.NumOfMeals,
			RightFork:  forks[(i+1)%info.NumOfPhilos],
			LeftFork:   forks[i%info.NumOfPhilos],
			Info:       info,
			LastEating: time.Now(),
		}
		i := i
		go func() {
			defer wg.Done()
			philos[i].lifeСircle(ctx)
		}()
	}
	endCh := make(chan bool)
	go checker(philos, endCh)
	for range endCh {
		<-endCh
	}
	cancel()
	wg.Wait()
}

func checkValid(info *Info) error {
	switch {
	case info.NumOfPhilos < 2:
		return errors.New("Need more philosophers!\n")
	case info.TimeToDie < 0:
		return errors.New("Negative value of time to die!\n")
	case info.TimeToEat < 0:
		return errors.New("Negative value of time to eat!\n")
	case info.TimeToSleep < 0:
		return errors.New("Negative value of time to sleep!\n")
	case info.NumOfMeals < 0:
		return errors.New("Negative value of meals!\n")
	default:
		return nil
	}
}

func philosophers(info *Info) {
	if err := checkValid(info); err != nil {
		log.Fatal(err)
		return
	}
	forks := createForks(info)
	info.start(forks)
	closeForks(forks)
	logger.Info("😎 ENDING SIMULATION 😎")
}
