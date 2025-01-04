package app

import (
	"context"
	"fmt"
	v1 "gapp1/api/v1"
	"gapp1/internal/config"
	"gapp1/internal/middleware"
	"gapp1/pkg/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// App 是整个应用的抽象
type App struct {
	server *http.Server
	wg     sync.WaitGroup
}

// NewApp 创建并初始化一个新的应用实例
func NewApp() *App {
	// 初始化 viper 配置
	config.LoadConfig()
	port := viper.GetString("server.port")

	db.ConnectDB()
	db.Migrate()

	// 创建 Gin 引擎
	r := gin.Default()

	r.Use(middleware.CORS())
	// r.Use(middleware.Auth())

	// 设置路由
	v1.SetupRoutes(r)

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &App{server: server}
}

// Run 启动应用
func (a *App) Run() {
	// 创建信号通道
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// 使用 WaitGroup 确保 goroutine 执行完成
	a.wg.Add(1)

	// 启动一个 goroutine 来监听信号
	go func() {
		defer a.wg.Done() // 保证退出时标记任务完成

		// 等待接收到信号
		sig := <-signalChan
		fmt.Printf("\nReceived signal: %s. Shutting down gracefully...\n", sig)

		// 创建带超时的上下文，设定 10 秒的停机时间
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// 调用 Stop 方法优雅停机
		a.Stop(ctx)

		db.CloseDB()
	}()

	// 启动 HTTP 服务
	fmt.Println("Starting server on", a.server.Addr)
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error starting server: %v\n", err)
	}

	// 等待 goroutine 完成
	a.wg.Wait()
}

// Stop 优雅关闭服务器
func (a *App) Stop(ctx context.Context) {
	log.Println("Attempting to shut down the server...")
	if err := a.server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	} else {
		fmt.Println("Server stopped gracefully.")
	}
}
