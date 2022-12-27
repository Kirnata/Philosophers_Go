package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func (philo *Philo) lifeСircle(ctx context.Context) {
	if philo.Id%2 == 1 {
		time.Sleep(5 * time.Millisecond)
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			philo.eating()
			philo.sleeping()
			philo.thinking()
		}
	}
}

func (philo *Philo) eating() {
	philo.LeftFork <- struct{}{}
	log.Info(fmt.Sprintf("Philo %d take left fork...", philo.Id))
	philo.RightFork <- struct{}{}
	log.Info(fmt.Sprintf("Philo %d take right fork...", philo.Id))
	log.Info(fmt.Sprintf("Philo %d is eating...", philo.Id))
	time.Sleep(philo.Info.TimeToEat)
	philo.LastEating = time.Now()
	philo.MealCount--
	<-philo.LeftFork
	<-philo.RightFork
}

func (philo *Philo) sleeping() {
	log.Info(fmt.Sprintf("Philo %d is sleeping...", philo.Id))
	time.Sleep(philo.Info.TimeToSleep)
}

func (philo *Philo) thinking() {
	log.Info(fmt.Sprintf("Philo %d is thinking...", philo.Id))
}
