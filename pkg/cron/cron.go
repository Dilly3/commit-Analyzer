package cron

import (
	"github.com/go-co-op/gocron"

	"time"
)

var schedular *gocron.Scheduler

// time in minutes (1 hour)
var timeDuration int64 = 30

// SetTimeDuration sets the time delay for the cron job
func SetTimeDuration(t int64) {
	timeDuration = t

}

// GetTimeDuration returns the time delay for the cron job
func GetTimeDuration() int64 {
	return timeDuration
}

// InitCron initializes the gocron scheduler
func InitCron() {

	s := gocron.NewScheduler(time.UTC)
	schedular = s
}

// StopCronJob stops the cron jobs
func StopCronJob() {
	schedular.Stop()
}

// StartCronJob starts the cron jobs
func StartCronJob() {
	schedular.StartAsync()
}

// SetCronJob sets the cron job
func SetCronJob(f func() error, interval int64) {
	timeD := time.Duration(interval) * time.Minute
	_, err := schedular.Every(timeD).Do(f)
	if err != nil {
		return
	}
}
