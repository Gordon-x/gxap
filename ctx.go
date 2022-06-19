//Package xap 内部公共代码
package xap

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var ctx *Ctx

func init() {
	ctx = &Ctx{}

	curDir, _ := os.Getwd()

	var configPath string
	flag.StringVar(&configPath, "config_path", filepath.Join(curDir, "config"), "config files path")
	flag.Parse()

	ctx.SysPath = sysPath{
		CurDir:    curDir,
		ConfigDir: configPath,
	}
	ctx.SysCtx = context.Background()

	ctx.InitSysConfig()
	ctx.InitLog()
	ctx.InitDb()
	ctx.InitRedis()
}

func GlobalContext() *Ctx {
	return ctx
}

func (ctx *Ctx) SignalListen(callback func(s os.Signal)) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigChan
	callback(sig)
}

func (ctx *Ctx) Close() {
	if ctx == nil {
		return
	}
	ctx.RedisClose()
}
