package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var logger = log.New()

func createForks(info *Info) []Fork {
	forks := make([]Fork, info.NumOfPhilos)
	for i := 0; i < info.NumOfPhilos; i++ {
		forks[i] = make(Fork, 1)
	}
	return forks
}

func closeForks(forks []Fork) {
	for _, v := range forks {
		close(v)
	}
}

func Start(philos []*Philo, forks []Fork, info *Info) {
	var wg sync.WaitGroup
	wg.Add(info.NumOfPhilos)
	logger.Info("Starting ...")

	ctx, cancel := context.WithCancel(context.Background())

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
			philos[i].lifecircle(ctx)
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

func philosophers(info *Info) {
	forks := createForks(info)
	philos := make([]*Philo, info.NumOfPhilos)
	Start(philos, forks, info)
	closeForks(forks)
	logger.Info("😎 ENDING SIMULATION 😎")
}
