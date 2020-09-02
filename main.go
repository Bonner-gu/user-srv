package main

import (
	"dm_user_srv/constant"

	"github.com/gogf/gf/frame/g"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name(constant.ServerName),
		micro.Address(constant.ServerAddress),
		micro.Version(constant.ServerVersion),
	)

	// Initialise service
	service.Init()

	// 注册服务
	userSrvPb.RegisterUserServiceHandler(service.Server(), new(handler.UserService))

	// Run service
	if err := service.Run(); err != nil {
		g.Log().Fatal(err)
	}
}
