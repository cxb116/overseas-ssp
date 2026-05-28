package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewClientMysql() *gorm.DB {
	if global.EngineConfig.Database.Username != "" {
		database := initMysqlDatabase(global.EngineConfig.Database)
		return database
	}
	return nil
}

func initMysqlDatabase(p config.DbBase) *gorm.DB {
	if p.Username == "" {
		return nil
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		p.Username, p.Password, p.Host, p.Port, p.DBname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键约束，提高性能
		DisableForeignKeyConstraintWhenMigrating: true,
		// 跳过默认事务
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Printf("mysql connection fail: %v", err)
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("mysql db handle fail: %v", err)
		return nil
	}
	// 优化连接池配置
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)      // 空闲连接数
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)      // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour)        // 连接最大存活时间1小时
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // 空闲连接最大存活时间10分钟

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		log.Printf("mysql ping fail: %v", err)
		_ = sqlDB.Close()
		return nil
	}

	log.Println("mysql connection success!")
	return db
}
