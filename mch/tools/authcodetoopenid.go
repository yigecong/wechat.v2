package tools

import (
	"github.com/chanxuehong/wechat.v2/mch/core"
)

// 授权码查询OPENID接口.
func AuthCodeToOpenId(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/tools/authcodetoopenid", req)
}
