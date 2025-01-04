package job

func (s *service) Stop() {
	if err := s.cron.Shutdown(); err != nil {
		s.logger.Error("stop cron service error: %v", err)
	}
}
