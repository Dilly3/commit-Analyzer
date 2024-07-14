package cron

import (
	"github.com/go-co-op/gocron"

	"time"
)

var schedular *gocron.Scheduler

// time in minutes
var timeDuration int64 = 3

func SetTimeDuration(t int64) {
	timeDuration = t

}
func GetTimeDuration() int64 {
	return timeDuration
}

func InitCron() {

	s := gocron.NewScheduler(time.UTC)
	schedular = s
}
func StopCronJob() {
	schedular.Stop()
}
func StartCronJob() {
	schedular.StartAsync()
}
func SetCronJob(f func() error, interval int64) {
	timeD := time.Duration(interval) * time.Minute
	_, err := schedular.Every(timeD).Do(f)
	if err != nil {
		return
	}
}
