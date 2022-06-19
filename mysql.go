package gxap

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// InitDb 连接数据库
func (ctx *Ctx) InitDb() {
	var err error
	dbLogger := logger.New(
		log.New(ctx.LogWriter("db.log"), "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second * 5,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		})
	conn, err := gorm.Open(mysql.Open(ctx.Config.Db.Dsn), &gorm.Config{
		AllowGlobalUpdate: false,
		CreateBatchSize:   1000,
		Logger:            dbLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   ctx.Config.Db.TablePrefix,
		},
	})
	if err != nil {
		ctx.Log.Error("mysql connect error %s", err.Error())
		panic(err)
	}
	if conn.Error != nil {
		ctx.Log.Error("database error %s", err.Error())
		panic(err)
	}

	db, _ := conn.DB()
	db.SetConnMaxIdleTime(time.Duration(ctx.Config.Db.ConnMaxLifeTime) * time.Second)
	db.SetMaxIdleConns(ctx.Config.Db.MaxIdleConn)
	db.SetMaxOpenConns(ctx.Config.Db.MaxOpenConn)
	ctx.Db = conn
}
