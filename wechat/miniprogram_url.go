package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//miniprogramUrl url
type miniprogramUrl struct {
	Config *MiniProgramlConfig
}

func newMiniprogramUrl(config *MiniProgramlConfig) miniprogramUrl {
	return miniprogramUrl{
		Config: config,
	}
}

// miniprogramGenerateschemeEntity 获取scheme码返回实体 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/queryScheme.html
type miniprogramGenerateschemeEntity struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Openlink string `json:"openlink"`
}

// GenerateScheme 获取用户基本信息 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/generateScheme.html
//	@param	param	map[string]interface	小程序code码
func (r *miniprogramUrl) GenerateScheme(param map[string]interface{}) miniprogramGenerateschemeEntity {
	var res miniprogramGenerateschemeEntity
	data, _ := json.Marshal(param)
	api := "https://api.weixin.qq.com/wxa/generatescheme?access_token=" + r.Config.AccessToken
	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(data)))
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
	}
	return res
}
