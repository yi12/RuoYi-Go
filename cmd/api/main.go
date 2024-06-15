// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package main

import (
	"RuoYi-Go/di"
	"RuoYi-Go/internal/shutdown"
	"RuoYi-Go/pkg/config"
	"os"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		os.Exit(2)
	}

	// 创建依赖注入容器
	container, err := di.NewContainer(cfg)
	if err != nil {
		os.Exit(2)
	}
	defer container.Close()

	//
	//app := iris.New()
	//ryserver.StartServer(app)
	//rywebsocket.StartWebSocket(app)
	//
	////log.Info("start server on:%d", conf.Server.Port)
	//app.Run(iris.Addr(fmt.Sprintf(":%d", cfg.Server.Port)))
	//
	//iris.RegisterOnInterrupt(func() {
	//	timeout := 5 * time.Second
	//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	//	defer cancel()
	//	// close all hosts
	//	app.Shutdown(ctx)
	//})

	// 系统关闭
	shutdown.NewHook().Close(
		func() {
			container.Close()
			os.Exit(0)
		},
	)
}
