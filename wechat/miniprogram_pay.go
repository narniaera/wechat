package wechat

import (
	"sort"
	"strings"
)

//miniprogramPay 支付
type miniprogramPay struct {
	Config *MiniProgramlConfig
}

func newMiniprogramPay(config *MiniProgramlConfig) miniprogramPay {
	return miniprogramPay{
		Config: config,
	}
}

//SignDataSort 签名参数排序
func (r miniprogramPay) SignDataSort(parmas map[string]string, key string) string {
	keys := []string{}
	for k, _ := range parmas {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str := ""
	for _, v := range keys {
		str += strings.Trim(parmas[v], "")
	}
	return str
}