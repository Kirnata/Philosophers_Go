package main

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func checker(philos []*Philo, endCh chan bool) {
	var mu sync.Mutex
	for {
		meal := 0
		for i := range philos {
			mu.Lock()
			deadline := philos[i].LastEating.Add(philos[i].Info.TimeToDie)
			if philos[i].MealCount == 0 {
				meal++
			}
			mu.Unlock()
			if time.Now().Sub(deadline) > 0 {
				log.Info("DEAD OF PHILO")
				log.Info("DEAD OF PHILO")
				time.Sleep(5 * time.Second)
				close(endCh)
				return
			}
			endCh <- true
		}
		if meal == len(philos) {
			close(endCh)
			return
		}
	}
}
