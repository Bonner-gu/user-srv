package handler

import (
	"context"
	"dm_user_srv/lib/validation"
	userSrvPb "dm_user_srv/proto/user"
	"dm_user_srv/service"
)

type UserService struct{}

//注册
func (u *UserService) UserSignup(ctx context.Context, req *userSrvPb.SignupReq, resp *userSrvPb.UserSignResp) error {
	//参数校验
	if err := validation.CkeckSignupParam(req); err != nil {
		return err
	}

	//service调用
	token, err := service.UserSrvCli.UserRegister(req.Phone, req.Passwd, req.Ip, req.Source)
	if err != nil {
		return err
	}

	resp.Token = token

	return nil
}

//登录
func (u *UserService) UserSignin(ctx context.Context, req *userSrvPb.SigninReq, resp *userSrvPb.UserSignResp) error {
	//参数校验
	if err := validation.CheckPhone(req.Phone); err != nil {
		return err
	}

	//service调用
	token, err := service.UserSrvCli.UserLoginCheck(req.Phone, req.Passwd)
	if err != nil {
		return err
	}

	resp.Token = token

	return nil
}
