package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//miniprogramCheck 内容安全审核
type miniprogramCheck struct {
	Config *MiniProgramlConfig
}

func newMiniprogramCheck(config *MiniProgramlConfig) miniprogramCheck {
	return miniprogramCheck{
		Config: config,
	}
}

//miniprogramCheckEntity 文本内容安全识别返回实体 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/msgSecCheck.html
type miniprogramCheckEntity struct {
	Errcode int                            `json:"errcode"`
	Errmsg  string                         `json:"errmsg"`
	TraceId string                         `json:"trace_id"`
	Result  miniprogramCheckResultEntity   `json:"result"`
	Detail  []miniprogramCheckDetailEntity `json:"detail"`
}

type miniprogramCheckResultEntity struct {
	Suggest string `json:"suggest"`
	Label   int64  `json:"label"`
}

type miniprogramCheckDetailEntity struct {
	Strategy string `json:"strategy"`
	Errcode  int    `json:"errcode"`
	Suggest  string `json:"suggest"`
	Label    int64  `json:"label"`
	Prob     int64  `json:"prob"`
	Level    int64  `json:"level"`
}

//Msg 文本审核 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/msgSecCheck.html
//https://developers.weixin.qq.com/miniprogram/dev/framework/security.imgSecCheck.html
//	@param	param	map[string]interface
func (r *miniprogramCheck) Msg(param map[string]interface{}) miniprogramCheckEntity {
	var res miniprogramCheckEntity
	data, _ := json.Marshal(param)
	api := "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=" + r.Config.AccessToken
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

//Media 媒体审核 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/mediaCheckAsync.html
//	@param	param	map[string]interface
func (r *miniprogramCheck) Media(param map[string]interface{}) miniprogramCheckEntity {
	var res miniprogramCheckEntity
	data, _ := json.Marshal(param)
	api := "https://api.weixin.qq.com/wxa/media_check_async?access_token=" + r.Config.AccessToken
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
