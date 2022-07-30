//Package gxap 内部公共代码
package gxap

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var ctx *Ctx

func Context() *Ctx {
	if ctx != nil {
		return ctx
	}
	ctx = &Ctx{}

	curDir, _ := os.Getwd()

	var configPath = filepath.Join(curDir, "config")
	ctx.SysPath = SysPath{
		CurDir:    curDir,
		ConfigDir: configPath,
	}
	ctx.SysCtx = context.Background()

	ctx.InitSysConfig()
	ctx.InitLog()
	ctx.InitDb()
	ctx.InitRedis()
	return ctx
}

func (ctx *Ctx) SignalListen(callback func(s os.Signal)) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case sig := <-sigChan:
		callback(sig)
	}
}

func (ctx *Ctx) Close() {
	if ctx == nil {
		return
	}
	ctx.RedisClose()
}

func (ctx *Ctx) GetConfig() Conf {
	return ctx.Config
}

func (ctx *Ctx) GetDb() *gorm.DB {
	return ctx.Db
}

func (ctx *Ctx) GetLog() *zap.SugaredLogger {
	return ctx.Log
}

func (ctx *Ctx) GetRedis() IRed {
	return ctx.Redis
}

func (ctx *Ctx) GetSysPath() SysPath {
	return ctx.SysPath
}
