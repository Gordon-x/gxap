package xap

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
)

type confDb struct {
	Dsn             string `yaml:"dsn"`
	ConnMaxLifeTime int    `yaml:"connMaxLifeTime"`
	MaxOpenConn     int    `yaml:"maxOpenConn"`
	MaxIdleConn     int    `yaml:"maxIdleConn"`
	TablePrefix     string `yaml:"tablePrefix"`
}

type confLog struct {
	Filepath   string `yaml:"filepath"`
	Filename   string `yaml:"filename"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"frequency"`
	MaxSize    int    `yaml:"maxSize"`
}

type confRedis struct {
	Type        string `yaml:"type"`
	Addr        string `yaml:"addr"`
	Password    string `yaml:"password"`
	Index       int    `yaml:"index"`
	MaxRedirect int    `yaml:"max-redirect"`
}

// Conf 配置文件格式
type Conf struct {
	Profiles string    `yaml:"profiles"`
	Log      confLog   `yaml:"log"`
	Db       confDb    `yaml:"db"`
	Redis    confRedis `yaml:"redis"`
}

type sysPath struct {
	CurDir    string
	ConfigDir string
}

type iRed interface {
	redis.Cmdable
	io.Closer
}

// Ctx 环境
type Ctx struct {
	Config  Conf
	Db      *gorm.DB
	Log     *zap.SugaredLogger
	SysPath sysPath
	Redis   iRed
	SysCtx  context.Context
}
