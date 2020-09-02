package service

var (
	UserSrvCli *UserSrv
)

func init() {
	UserSrvCli = NewUserSrv()
}
