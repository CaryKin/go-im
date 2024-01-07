package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go_im/lib/mysqllib"
	"go_im/lib/redislib"
	"go_im/routers"
	"io"
	"net/http"
	"os"
)

func main() {
	initConfig()

	initFile()

	initRedis()

	initMysql()

	router := gin.Default()
	// 初始化路由
	routers.Init(router)

	httpPort := viper.GetString("app.httpPort")
	http.ListenAndServe(":"+httpPort, router)
}

// 初始化配置
func initConfig() {
	viper.SetConfigName("config/app")
	viper.AddConfigPath(".") // 添加搜索路径

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config redis:", viper.Get("redis"))
}

// 初始化日志
func initFile() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	logFile := viper.GetString("app.logFile")
	f, _ := os.Create(logFile)
	gin.DefaultWriter = io.MultiWriter(f)
}

// 初始化redis
func initRedis() {
	redislib.ExampleNewClient()
}

// 初始化mysql
func initMysql() {
	mysqllib.ExampleNewClient()
}
