package main

import (
	"context"
	"errors"
	"flag"
	"hook007/config"
	"hook007/dao/query"
	"hook007/job"
	"hook007/model"
	"hook007/pkg/cache"
	"hook007/pkg/db"
	"hook007/pkg/utils"
	"hook007/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	withCron   = flag.Bool("with-cron", false, "run with cron service")
	configPath = flag.String("conf", "./", "set config path")
)

func main() {
	flag.Parse()
	config.ConfigPath = *configPath
	config.Init()

	utils.InitSlog()

	mysqlDB, err := db.Connect()
	if err != nil {
		panic(err)
	}
	model.AutoMigrate(mysqlDB)
	query.SetDefault(mysqlDB)

	redisCache, err := cache.Connect()
	if err != nil {
		panic(err)
	}

	cronTask := job.New()
	if *withCron {
		cronTask.Start()
	}

	s, err := router.NewRouter(redisCache, query.Q)
	if err != nil {
		panic(err)
	}

	httpServer := &http.Server{
		Addr:    ":" + config.Get().Server.Port,
		Handler: s.Mux,
	}

	go func() {
		ip, err := utils.GetLocalIP()
		if err != nil {
			log.Fatal("GetLocalIP error:", err)
		}
		log.Printf("Server is running at: %s\n", ip+httpServer.Addr)

		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Server error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	sqlDB, err := mysqlDB.DB()
	if err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("db connection close err:", err)
		}
	}

	if redisCache != nil {
		if err := redisCache.Close(); err != nil {
			log.Fatal("cache connection close err:", err)
		}
	}

	if *withCron {
		cronTask.Stop()
	}

	log.Println("Server exiting")
}
