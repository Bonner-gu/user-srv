package constant

import "github.com/gogf/gf/frame/g"

var (
	ServerName    = g.Config().GetString("app.ServerName")
	ServerAddress = g.Config().GetString("app.Address")
	ServerVersion = g.Config().GetString("app.Version")
	Cli4suer      = g.Config().GetString("app.Cli4user")
)

var (
	SrvCli4RegAddr = g.Config().GetString("srvcli.SrvCli4RegAddr")
	SrvCli4BaseCli = g.Config().GetString("srvcli.SrvCli4BaseCli")
	SrvCli4AuthSrv = g.Config().GetString("srvcli.SrvCli4AuthSrv")
)

// base
const (
	BaseId   = 10000 // 默认id
	BasePage = 1     // 默认页码
	BaseSize = 20    // 默认每页展示数量
)
