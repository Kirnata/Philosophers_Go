package main

import (
	"context"
	"time"
)

func (philo *Philo) lifecircle(ctx context.Context) {
	if philo.Id%2 == 1 {
		time.Sleep(5 * time.Millisecond)
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		philo.eating()
		philo.sleeping()
		philo.thinking()
	}
}

func (philo *Philo) eating() {
	philo.LeftFork <- struct{}{}
	logger.Info("Take left fork...")
	philo.RightFork <- struct{}{}
	logger.Info("Take right fork...")
	logger.Info("Eating...")
	time.Sleep(philo.Info.TimeToEat)
	philo.LastEating = time.Now()
	philo.MealCount--
	<-philo.LeftFork
	<-philo.RightFork
}

func (philo *Philo) sleeping() {
	logger.Info("Sleaping...")
	time.Sleep(philo.Info.TimeToSleep)
}

func (philo *Philo) thinking() {
	logger.Info("Thinking...")
}
