package main

import (
	"fmt"
	"github.com/narniaera/wechat/entity"
	"github.com/narniaera/wechat/wechat"
)

func main() {
	ofc()
	ofv()
}

var res entity.ResultEntity

func ofc() {
	official := wechat.NewOfficial(&wechat.OfficialConfig{
		Appid:     "wxc8399122cea914bc",
		Appsercet: "ab47cc15caaffb9fcc7326113f66d629",
		Token:     "token",
	})
	//获取access_token  56_c3_0Uypdb9HDIejT0AN8CeZaHUlZY5mN_XyNDqasjPx1zo5z6zuflzG6jkJk0ptrDm-040uzvO_BoNZh2cDBqOIap1zusfMvDQGlASLrSYcSQotHtoYxWtvR1RLzPIqK66Afey9jFFnKkzYxFSGhAIAJWK
	//official.GetAccessToken(true)
	official.SetAccessToken("56_dARMoTtfKWqy0_TMRDRJTUw7vHkd_xlA0o5KNKZmkVQLejzitXx6USe8lOAb-51rwoVy--I4po4pJrgYQ0xdsil_vhE6cLY66UAsFrrarkwH5H_LkZ12sGFOYxmitGsnO_S8fOExejNqoIN0CGRaAAATXI")
	//fmt.Println("access_token:" + official.Config.AccessToken)
	//
	//res := official.User.GetUserList("")
	//fmt.Println(res)
	//b, _ := json.Marshal(res)
	//fmt.Println(string(b))

	//var r1 entity.ResultEntity
	//res.Errcode = 1
	fmt.Println(res)
}

func ofv() {
	//var res entity.ResultEntity
	//res.Errmsg = "aaa"
	//fmt.Println(res)
}
