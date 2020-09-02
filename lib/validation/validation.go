package validation

import (
	userSrvPb "dm_user_srv/proto/user"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gvalid"
)

//注册校验
func CkeckSignupParam(req *userSrvPb.SignupReq) error {
	params := map[string]interface{}{
		"ip":      req.Ip,
		"phone":   req.Phone,
		"newpw":   req.Passwd,
		"rightpw": req.RightPasswd,
		"source":  req.Source,
	}
	rules := map[string]string{
		"ip":      "required|ipv4",
		"phone":   "required|phone",
		"newpw":   "required|same:rightpw",
		"rightpw": "required",
		"source":  "required|length:1,8",
	}
	msgs := map[string]interface{}{
		"ip":      "ip不能为空|ip格式不合法",
		"account": "账号不能为空|手机号格式不合法",
		"newpw": map[string]string{
			"required": "密码不能为空",
			"same":     "两次密码不一致",
		},
		"rightpw": map[string]string{
			"required": "密码不能为空",
		},
		"source": "来源不能为空|来源格式不合法",
	}
	if err := gvalid.CheckMap(params, rules, msgs); err != nil {
		g.Log().Errorf("注册参数校验失败 [params:%v]", params)
		return err
	} else {
		return nil
	}
}

//密码校验
func CheckPw(pw, rightpw string) error {
	params := map[string]interface{}{
		"pw":      pw,
		"rightpw": rightpw,
	}
	rules := map[string]string{
		"pw":      "required|length:6,16|same:rightpw",
		"rightpw": "required|length:6,16",
	}
	msgs := map[string]interface{}{
		"pw": map[string]string{
			"required": "密码不能为空",
			"same":     "两次密码不一致",
			"length":   "密码格式不合法，密码长度应在:min ~ :max之间",
		},
		"rightpw": map[string]string{
			"required": "密码不能为空",
			"length":   "密码格式不合法，密码长度应在:min ~ :max之间",
		},
	}
	if err := gvalid.CheckMap(params, rules, msgs); err != nil {
		g.Log().Errorf("密码校验失败 [params:%v]", params)
		return err
	} else {
		return nil
	}
}

//手机号码校验
func CheckPhone(phone string) error {
	if err := gvalid.Check(phone, "required|phone", "手机号码必填|手机号码格式不合法"); err != nil {
		g.Log().Errorf("phone格式校验失败 [phone:%v]", phone)
		return err
	} else {
		return nil
	}
}

//来源校验
func CheckSource(source string) error {
	if err := gvalid.Check(source, "required|length:1,8", "来源不能为空|来源格式不合法"); err != nil {
		g.Log().Errorf("source格式校验失败 [source:%v]", source)
		return err
	} else {
		return nil
	}
}

//IP校验
func CheckIp(ip string) error {
	if err := gvalid.Check(ip, "required|ipv4", "ip不能为空|ip格式不合法"); err != nil {
		g.Log().Errorf("ip格式校验失败 [ip:%v]", ip)
		return err
	} else {
		return nil
	}
}

//手机号*号保密处理
func PhoneHide(phone string) string {
	if len(phone) <= 10 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-4:]
}

//身份证*号保密处理
func IdCardHide(card string) string {
	if len(card) <= 17 {
		return card
	}
	return card[:6] + "********" + card[len(card)-4:]
}
