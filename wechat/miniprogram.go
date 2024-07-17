package wechat

import (
	"bytes"
	"encoding/json"
	"github.com/narniaera/wechat/entity"
	"github.com/narniaera/wechat/utils"
	"io/ioutil"
	"net/http"
	"sort"
)

// MiniProgramlConfig 配置
type MiniProgramlConfig struct {
	Appid         string
	Appsercet     string
	Token         string
	AccessToken   string
	Request       *http.Request
	ResponseWrite http.ResponseWriter
}

// Official 小程序服务类
type MiniProgram struct {
	Config  *MiniProgramlConfig
	User    miniprogramUser
	Url     miniprogramUrl
	Check   miniprogramCheck
	Vod     miniprogramVod
	Message miniprogramMessage
}

// NewOfficial 新建公众号
func NewMiniProgram(config *MiniProgramlConfig) MiniProgram {
	return MiniProgram{
		Config:  config,
		User:    newMiniprogramUser(config),
		Url:     newMiniprogramUrl(config),
		Check:   newMiniprogramCheck(config),
		Vod:     newMiniprogramVod(config),
		Message: newMiniprogramMessage(config),
	}
}

//GetAccessToken 获取access_token
//	@return	string token | ""
func (r *MiniProgram) GetAccessToken(b bool) string {
	if b {
		var res map[string]interface{}
		api := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + r.Config.Appid + "&secret=" + r.Config.Appsercet
		resp, err := http.Get(api)
		defer resp.Body.Close()
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &res)
			if res["errcode"] == nil {
				r.Config.AccessToken = res["access_token"].(string)
			} else {
				r.Config.AccessToken = ""
			}
		}
	}
	return r.Config.AccessToken
}

//GetStableAccessToken 稳定版获取access_token https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getStableAccessToken.html
//	@param force bool 是否强制刷新
func (r *MiniProgram) GetStableAccessToken(force bool) entity.ResultAccessTokenEntity {
	var res entity.ResultAccessTokenEntity
	request := map[string]string{
		"grant_type": "client_credential",
		"appid":      r.Config.Appid,
		"secret":     r.Config.Appsercet,
	}
	api := "https://api.weixin.qq.com/cgi-bin/stable_token"
	data, _ := json.Marshal(request)
	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(data)))
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
		if res.Errcode == 0 {
			r.Config.AccessToken = res.AccessToken
		} else {
			r.Config.AccessToken = ""
		}
	}
	return res
}

// SetAccessToken 设置access_token
//	@param accessToken string	access_token
func (r *MiniProgram) SetAccessToken(accessToken string) {
	r.Config.AccessToken = accessToken
	r.Config.Token = accessToken
}

// Send 服务推送
func (r *MiniProgram) Send() {
	signature := r.Config.Request.URL.Query().Get("signature")
	nonce := r.Config.Request.URL.Query().Get("nonce")
	timestamp := r.Config.Request.URL.Query().Get("timestamp")
	token := "token"
	echostr := r.Config.Request.URL.Query().Get("echostr")
	tmps := []string{token, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	tmp := utils.WxSha1(tmpStr)
	if tmp == signature {
		r.Config.ResponseWrite.Write([]byte(string(echostr)))
	}
}

// Handle 消息监听
// 	@return OfficialMessage
func (r *MiniProgram) Handle(callback func(entity entity.MiniprogramMessageAcceptEntity)) {
	body, err := ioutil.ReadAll(r.Config.Request.Body)
	messageEntity := entity.MiniprogramMessageAcceptEntity{}
	if err != nil {
		//log.Println(err.Error())
	} else {
		defer r.Config.Request.Body.Close()
		json.Unmarshal(body, &messageEntity)
	}
	if messageEntity.FromUserName != "" {
		callback(messageEntity)
	}
}

//// Qrcode https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
//// 	@param	data map[string]interface{}
////	@return	string	二维码地址
//func (r *Official) Qrcode(data map[string]interface{}) string {
//	var qrcode string
//	param, _ := json.Marshal(data)
//	api := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + r.Config.AccessToken
//	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(param)))
//	defer resp.Body.Close()
//	if err != nil {
//		log.Println(err.Error())
//	} else {
//		body, _ := ioutil.ReadAll(resp.Body)
//		var ticket map[string]interface{}
//		json.Unmarshal(body, &ticket)
//		if ticket["ticket"] != nil {
//			qrcode = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(ticket["ticket"].(string))
//		}
//	}
//	return qrcode
//}
