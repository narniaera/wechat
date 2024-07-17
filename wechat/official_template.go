package wechat

import (
	"bytes"
	"encoding/json"
	"github.com/narniaera/wechat/entity"
	"io/ioutil"
	"log"
	"net/http"
)

// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.htm
type officialTemplate struct {
	Config *OfficialConfig
}

func newOfficialTemplate(config *OfficialConfig) officialTemplate {
	return officialTemplate{
		Config: config,
	}
}

// 获取模板消息列表实体
type getAllTemplateEntity struct {
	TemplateList []struct {
		TemplateId     string `json:"template_id"`
		Title          string `json:"title"`
		PrimaryIndusty string `json:"primary_industy"`
		DeputyIndustry string `json:"deputy_industry"`
		Content        string `json:"content"`
		Example        string `json:"example"`
	} `json:"template_list"`
}

// GetAllTemplateList 获取模板消息列表 https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.htm
func (r *officialTemplate) GetAllTemplateList() getAllTemplateEntity {
	var res getAllTemplateEntity
	api := "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=" + r.Config.AccessToken
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

// Delete 删除模板消息 https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.htm
func (r *officialTemplate) Delete(data map[string]string) entity.ResultEntity {
	var res entity.ResultEntity
	param, _ := json.Marshal(data)
	api := "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=" + r.Config.AccessToken
	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(param)))
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
	}
	return res
}

//模板消息发送实体
type templateSendEntity struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	MsgId   int64  `json:"msg_id"`
}

// Send 发送模板消息 https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.htm
func (r *officialTemplate) Send(data map[string]interface{}) map[string]interface{} {
	var res map[string]interface{}
	param, _ := json.Marshal(data)
	api := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + r.Config.AccessToken
	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(param)))
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
	}
	return res
}
