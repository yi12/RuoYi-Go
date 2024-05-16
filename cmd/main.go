// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package main

import (
	"RuoYi-Go/internal/server"
	"RuoYi-Go/internal/shutdown"
	ws "RuoYi-Go/internal/websocket"
	"RuoYi-Go/pkg/config"
	"RuoYi-Go/pkg/db"
	"RuoYi-Go/pkg/i18n"
	"RuoYi-Go/pkg/logger"
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"time"

	"context"
)

func main() {
	// 初始化配置
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	// 初始化日志
	log, err := logger.InitializeLogger(conf.Debug) // 假设配置中有Debug字段
	if err != nil {
		panic(err)
	}

	// 初始化国际化
	_, err = i18n.GetLocalizer(conf.Language) // 假设配置中指定了Language
	if err != nil {
		log.Error("Failed to get localizer", zap.Error(err))
	}

	// 创建DatabaseStruct实例
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Shanghai", conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DBName)
	sqlService := &db.DatabaseStruct{
		DatabaseType: conf.Database.DBtype,
		Dsn:          dsn, // SQLite 数据库文件路径
	}
	sqlService.OpenSqlite()

	app := iris.New()

	server.InitServer(app)
	server.StartServer()
	ws.InitWebSocket(app)
	ws.StartWebSocket()

	app.Run(iris.Addr(fmt.Sprintf(":%d", conf.Server.Port)))

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})

	// 系统关闭
	shutdown.NewHook().Close(
		// 关闭 logger
		func() {
			logger.Close()
		},

		func() {
			err = sqlService.CloseSqlite()
			if err != nil {
				log.Error("Failed to close the database connection:", zap.Error(err))
			}
		},
	)
}
