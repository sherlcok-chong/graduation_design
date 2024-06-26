package main

import (
	v1 "GraduationDesign/src/api/v1"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"GraduationDesign/src/global"
	"GraduationDesign/src/routing/router"
	"GraduationDesign/src/setting"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy

func init() {
	log.SetFlags(log.Ltime | log.Llongfile)
}

// @title        chat
// @version      1.0
// @description  在线聊天系统

// @license.name  raja,chong
// @license.url

// @host      chat.humraja.xyz
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {
	setting.AllInit()
	if global.PbSettings.Server.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.NewRouter() // 注册路由
	s := &http.Server{
		Addr:           global.PbSettings.Server.Address,
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Info("Server started!")
	fmt.Println("AppName:", global.PbSettings.App.Name, "Version:", global.PbSettings.App.Version, "Address:", global.PbSettings.Server.Address, "RunMode:", global.PbSettings.Server.RunMode)
	errChan := make(chan error, 1)
	defer close(errChan)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()
	go v1.Group.Ws.Broadcast()
	// 优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errChan:
		global.Logger.Error(err.Error())
	case <-quit:
		global.Logger.Info("ShutDown Server...")
		// 给几秒完成剩余任务
		ctx, cancel := context.WithTimeout(context.Background(), global.PbSettings.Server.DefaultContextTimeout)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil { // 优雅退出
			if !errors.Is(err, context.DeadlineExceeded) {
				global.Logger.Error("Server forced to ShutDown,Err:" + err.Error())
			}
		}
	}
	global.Logger.Info("Server exited!")
}
