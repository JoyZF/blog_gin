package main

import (
	"flag"
	"github.com/JoyZF/blog_gin/global"
	"github.com/JoyZF/blog_gin/internal/model"
	"github.com/JoyZF/blog_gin/internal/routers"
	"github.com/JoyZF/blog_gin/pkg/logger"
	"github.com/JoyZF/blog_gin/pkg/setting"
	"github.com/JoyZF/blog_gin/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	port    string
	runMode string
	config  string
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err : %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDbEngine err : %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLoger err : %v", err)
	}
	err = setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err : %v",err)
	}
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof( "监听端口%s", global.ServerSetting.HttpPort)
	err := s.ListenAndServe()
	if err != nil {
		global.Logger.Panicf( "项目启动失败", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting(strings.Split(config,",")...)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	global.JWTSetting.Expire *= time.Second

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-gin",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}

func setupFlag() error {
	flag.StringVar(&port,"port","","启动端口")
	flag.StringVar(&runMode,"mode","","启动模式")
	flag.StringVar(&config,"config","configs/","指定要使用的配置文件路径")
	flag.Parse()
	return nil
}
