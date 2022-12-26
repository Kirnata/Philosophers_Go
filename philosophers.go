package main

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var logger = log.New()

func CreateForks(info *Info) []Fork {
	forks := make([]Fork, info.NumOfPhilos)
	for i := 0; i < info.NumOfPhilos; i++ {
		forks[i] = make(Fork)
	}
	return forks
}

func Start(philos []*Philo, forks []Fork, info *Info) {
	var wg sync.WaitGroup
	wg.Add(info.NumOfPhilos)
	logger.Info("Starting ...")

	for i := range philos {
		philos[i] = &Philo{
			Id:        i + 1,
			MealCount: info.NumOfMeals,
			RightFork: forks[(i+1)%info.NumOfPhilos],
			LeftFork:  forks[i%info.NumOfPhilos],
			Info:      info,
		}
		i := i
		go func() {
			defer wg.Done()
			Lifecircle(philos[i])
		}()
	}

	wg.Wait()
}

func Lifecircle(philo *Philo) {
	if philo.Id%2 == 1 {
		time.Sleep(5 * time.Millisecond)
	}
	for philo.MealCount > 0 {
		logger.Info("Eating...")
		time.Sleep(philo.Info.TimeToEat)
		logger.Info("Sleaping...")
		time.Sleep(philo.Info.TimeToSleep)
		logger.Info("Thinking...")
		logger.Info("Thinking...")
		philo.MealCount--
	}
}

func philosophers(info *Info) {
	forks := CreateForks(info)
	philos := make([]*Philo, info.NumOfPhilos)
	//makeForks(forks)
	Start(philos, forks, info)
	logger.Info("😎 ENDING SIMULATION 😎")
}
