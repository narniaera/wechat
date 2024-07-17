package wechat

import (
	"bytes"
	"encoding/json"
	"github.com/narniaera/wechat/entity"
	"io/ioutil"
	"log"
	"net/http"
)

//miniprogramMessage 消息
type miniprogramMessage struct {
	Config *MiniProgramlConfig
}

func newMiniprogramMessage(config *MiniProgramlConfig) miniprogramMessage {
	return miniprogramMessage{
		Config: config,
	}
}

//模板
type miniprogramTemplateEntity struct {
	Errcode int                             `json:"errcode"`
	Errmsg  string                          `json:"errmsg"`
	Data    []miniprogramTemplateDataEntity `json:"data"`
}

//模板数据
type miniprogramTemplateDataEntity struct {
	PriTmplId            string `json:"priTmplId"`
	Title                string `json:"title"`
	Content              string `json:"content"`
	Example              string `json:"example"`
	Type                 int    `json:"type"`
	KeywordEnumValueList []miniprogramTemplateDataKeywordEnumValueListEntity `json:"keywordEnumValueList"`
}

//模板数据枚举参数值范围
type miniprogramTemplateDataKeywordEnumValueListEntity struct {
	EnumValueList []string `json:"enumValueList"`
	KeywordCode   string   `json:"keywordCode"`
}

//SubscribeSend 发送订阅消息 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/sendMessage.html
//	@param	param	map[string]interface
func (r *miniprogramMessage) SubscribeSend(param map[string]interface{}) entity.ResultEntity {
	var res entity.ResultEntity
	data, _ := json.Marshal(param)
	api := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + r.Config.AccessToken
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

//GetTemplate 获取订阅消息模板 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/getMessageTemplateList.html
//	@param	param	map[string]interface
func (r *miniprogramMessage) GetTemplate() miniprogramTemplateEntity {
	var res miniprogramTemplateEntity
	api := "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token=" + r.Config.AccessToken
	resp, err := http.Get(api)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
	}
	return res
}
