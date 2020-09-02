package cli

import (
	"dm_user_srv/constant"

	"github.com/gogf/gf/util/gconv"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var (
//AuthSrvCli authSrvPb.AuthService
)

func init() {
	// 启用注册中心
	req := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = gconv.Strings(constant.SrvCli4RegAddr)
	})

	// 创建客户端服务
	srv := micro.NewService(
		micro.Registry(req),
		micro.Name(constant.SrvCli4BaseCli),
	)

	// 初始化
	srv.Init()

	// 创建客户端
	//AuthSrvCli = authSrvPb.NewAuthService(constant.SrvCli4AuthSrv, srv.Client())
}
