package service

import (
	"dm_user_srv/lib/MD5"
	"dm_user_srv/lib/ecode"
	"dm_user_srv/model/t_user"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type UserSrv struct{}

func NewUserSrv() *UserSrv {
	return &UserSrv{}
}

//用户注册
func (u *UserSrv) UserRegister(phone, passwd, ip, source string) (token string, err error) {
	//判断用户账号是否存在
	user := t_user.GetOne(g.Map{"phone": phone})
	if user != nil {
		g.Log().Line().Debugf("The user already exists！")
		err = ecode.UserExistErr
	}
	g.Log().Line().Debugf("This account can be used， Start registration！")

	//时间戳+用户手机号加密生成动态盐
	salt := MD5.Md5Str(gconv.String(gtime.Timestamp()) + phone)

	//动态盐+密码生成最终密码串
	pwStr := MD5.Md5Str(salt + passwd)

	//DB库操作
	userInfo := &t_user.Entity{
		Nickname:   "user_name",
		Phone:      phone,
		Salt:       salt,
		Passwd:     pwStr,
		Logo:       "logo.jpg",
		Source:     "PC",
		RegisterIp: ip,
		CreateTime: gtime.Now(),
	}
	if err = t_user.CreateOne(userInfo); err != nil {
		g.Log().Line().Errorf("DB database operation failed：[error:%v]", err)
		err = ecode.RegisterFailed
		return
	}

	//生成token

	token = ""

	return
}

//登录验证
func (u *UserSrv) UserLoginCheck(phone, passwd string) (token string, err error) {
	//获取用户数据
	user := t_user.GetOne(g.Map{"phone": phone})
	if user == nil {
		g.Log().Line().Errorf("User does not exist！")
		err = ecode.NotLogin
		return
	}

	//输入的密码和动态盐生成最终密码串与DB库密码串校验
	pw := MD5.Md5Str(user.Salt + passwd)
	if user.Passwd != pw {
		g.Log().Line().Errorf("The two passwords do not match！")
		err = ecode.PasswordIncorrect
		return
	}

	//生成token

	token = ""

	return
}
