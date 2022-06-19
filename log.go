package gxap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"path"
)

func (ctx *Ctx) InitLog() {
	err := CreatePath(ctx.Config.Log.Filepath)
	if err != nil {
		panic(err)
	}

	encoder := ctx.getEncoder()
	writeSyncer := ctx.getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())
	ctx.Log = logger.Sugar()
}

func (ctx *Ctx) getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (ctx *Ctx) getLogWriter() zapcore.WriteSyncer {
	logConf := ctx.Config.Log
	lumberJackLogger := ctx.LogWriter(logConf.Filename)
	return zapcore.AddSync(lumberJackLogger)
}

func (ctx *Ctx) LogWriter(name string) io.Writer {
	logConf := ctx.Config.Log
	return &lumberjack.Logger{
		Filename:   path.Join(logConf.Filepath, name),
		MaxSize:    logConf.MaxSize,
		MaxBackups: logConf.MaxBackups,
		MaxAge:     logConf.MaxAge,
		Compress:   false,
	}
}
