package job

import (
	"github.com/go-co-op/gocron/v2"
)

type (
	Service interface {
		Start()
		Stop()
	}

	service struct {
		cron   gocron.Scheduler
		logger gocron.Logger
	}
)

func New() Service {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	return &service{
		cron:   s,
		logger: gocron.NewLogger(gocron.LogLevelDebug),
	}
}
