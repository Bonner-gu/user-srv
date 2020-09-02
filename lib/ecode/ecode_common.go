package ecode

var (
	//访问成功 = 0
	OK        = add(0)
	ServerErr = add(-1000)

	NoPermission = add(2002) //没有权限

	// 系统公共错误 < 0
	MysqlManageErr     = add(-603)  // db操作失败
	Md5ManageErr       = add(-604)  // md5操作失败
	AccessDenied       = add(-403)  // 访问权限不足
	AuthParamErr       = add(8001)  // 参数错误
	AuthTokenMakeErr   = add(8002)  // token生成错误
	AuthTokenParseErr  = add(8003)  // token解析错误
	AuthTokenFoundErr  = add(8004)  // token创建错误
	NotLogin           = add(-2001) //登录失败
	RegisterFailed     = add(-2002) //注册失败
	PasswordIncorrect  = add(-2003) //密码验证错误
	UserExistErr       = add(-2004) //注册失败用户已存在
	VercodeMakeErr     = add(-2005) //手机验证码生成错误
	VercodeCheckErr    = add(-2006) //手机验证码验证错误
	VercodeCheckFailed = add(-2007) //手机验证码验证
)
