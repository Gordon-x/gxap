// Package gxap 项目内公用工具类
package gxap

import (
	"github.com/go-redis/redis/v8"
	"strings"
)

// InitRedis 连接redis
func (ctx *Ctx) InitRedis() {
	rc := ctx.Config.Redis

	if rc.Type == "cluster" {
		ctx.Redis = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        strings.Split(rc.Addr, ","),
			Password:     rc.Password,
			MaxRedirects: 6,
		})
	} else {
		ctx.Redis = redis.NewClient(&redis.Options{
			Addr:     rc.Addr,
			Password: rc.Password, // no password set
			DB:       rc.Index,    // use default DB
		})
	}
}

// RedisClose 关闭redis客户端
func (ctx *Ctx) RedisClose() {
	if ctx.Redis == nil {
		return
	}
	ctx.Log.Info("Redis连接关闭")
	err := ctx.Redis.Close()
	if err != nil {
		ctx.Log.Error(err.Error())
	}
}
