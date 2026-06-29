package main

import (
	"goCode/go_blog/config"
	"goCode/go_blog/routes"
	"log"
)

func main() {

	// 初始化日志
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//logrus.Info("Starting blog application...")

	// 初始化数据库链接
	config.InitDatabase()

	// 设置路由
	r := routes.SetupRoutes()

	// 启动服务器
	port := config.SysConfig.Server.Port
	port = ":" + port
	if err := r.Run(port); err != nil {
		log.Fatal("服务启动失败:", err)
	}

	log.Print("服务启动成功")
}
