package initialize

import (
	"bankroll/global"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	router := Routers()
	s := initServer(address, router)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}