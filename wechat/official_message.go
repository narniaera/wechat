package wechat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/narniaera/wechat/entity"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type officialMessage struct {
	Config *OfficialConfig
}

func newOfficialMessage(config *OfficialConfig) officialMessage {
	return officialMessage{
		Config: config,
	}
}

// Push 推送消息	https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Passive_user_reply_message.html
func (r *officialMessage) Push(messageEntity entity.OfficialAccountMessageEntity) {
	messageEntity.CreateTime = time.Now().Unix()
	msg, _ := xml.Marshal(messageEntity)
	r.Config.ResponseWrite.Write(msg)
}

// Custom 发送客服消息 https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html#%E5%AE%A2%E6%9C%8D%E6%8E%A5%E5%8F%A3-%E5%8F%91%E6%B6%88%E6%81%AF
func (r *officialMessage) Custom(data map[string]interface{}) entity.ResultEntity {
	var res entity.ResultEntity
	param, _ := json.Marshal(data)
	api := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + r.Config.AccessToken
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
