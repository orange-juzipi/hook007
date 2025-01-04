package job

func (s *service) Start() {
	// logger := s.logger

	// s.cron.NewJob(
	// 	gocron.CronJob("0/3 * * * * ?", true),
	// 	// 有参数
	// 	// gocron.NewTask(func(a string) {
	// 	// 	logger.Info("cron job run: %v", a)
	// 	// }, "a"),

	// 	// 无参数
	// 	gocron.NewTask(func() {
	// 		logger.Info("cron job run")
	// 	}),
	// )

	s.cron.Start()
}
