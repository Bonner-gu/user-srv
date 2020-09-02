package MD5

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
)

func Md5Str(val string) string {
	res, err := gmd5.Encrypt(val)
	if err != nil {
		g.Log().Line().Errorf("MD5 encryption failed!")
		return ""
	}
	return res
}
