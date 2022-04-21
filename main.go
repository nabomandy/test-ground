package main

import (
	"context"
	"fmt"

	// "encoding/json"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/config"
	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/dao"
	router "dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/process"

	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/process/api/v1/dbData"
	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/process/seeds"

	// "dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/process/sync"
	//"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error { return nil }

func InitLogger() {
	zap.RegisterSink("lumberjack", func(u *url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &lumberjack.Logger{
				Filename:   u.Opaque,
				MaxSize:    100, // megabytes
				MaxBackups: 10,
				MaxAge:     30,   // days
				Compress:   true, // disabled by default
			},
		}, nil
	})

	logPath := "./logs"
	logFileName := "app.log"

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.Mkdir(logPath, os.ModePerm)
	}
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"stderr",
		logPath + "/" + logFileName,
	}
	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)
}

// func init() {
// 	testv := dbData.GetDataList2()
// 	helloworld = testv[0].Prop_value
// }

// testing
func main() {
	// 설정 로딩
	config.GetInstance().LoadConfig() //jwj monster

	// 로그서비스 초기화
	InitLogger()
	// DB Migration & Seeding
	//InitDB()
	testv := dbData.GetDataList2()
	fmt.Println(testv[0].Prop_value)

	// API 라우터 초기화
	r := router.InitRouter()

	srv := &http.Server{
		Addr:    ":" + cast.ToString(config.GetInstance().Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()

	zap.S().Infof("Starting server... port : %v", config.GetInstance().Port)

	// 킬 시그널 (ctrl+c) 처리부분
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("Shutting down server...")

	// 스케줄링 워커 종료
	jobctx := dao.CronScheduler.Stop()
	<-jobctx.Done()

	// go-gin 서비스 종료 (타임아웃 10초)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Fatalf("Server forced to shutdown: %v", err)
	}

	zap.S().Info("Server exiting")

}

func InitDB() {
	seeder := seeds.NewSeeder(config.GetInstance().VurixDmsDB)
	seeder.CreateDemoProgramSeeder()
	seeder.CheckExists()

}
