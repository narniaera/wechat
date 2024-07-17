package wechat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/narniaera/wechat/entity"
	"github.com/narniaera/wechat/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
)

// OfficialConfig 配置
type OfficialConfig struct {
	Appid         string
	Appsercet     string
	Token         string
	AccessToken   string
	Request       *http.Request
	ResponseWrite http.ResponseWriter
}

// Official 公众号服务类
type Official struct {
	Config   *OfficialConfig
	Menu     officialMenu
	Message  officialMessage
	Template officialTemplate
	User     officialUser
	Material officialMaterial
}

// NewOfficial 新建公众号
func NewOfficial(config *OfficialConfig) Official {
	return Official{
		Config:   config,
		Menu:     newOfficialMenu(config),
		Message:  newOfficialMessage(config),
		Template: newOfficialTemplate(config),
		User:     newOfficialUser(config),
		Material: newOfficialMaterial(config),
	}
}

// GetAccessToken 获取access_token
//	@param b bool	是否刷新获取
//	@return	string token | ""
func (r *Official) GetAccessToken(b bool) string {
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
func (r *Official) GetStableAccessToken(force bool) entity.ResultAccessTokenEntity {
	var res entity.ResultAccessTokenEntity
	request := map[string]string{
		"grant_type": "client_credential",
		"appid":      r.Config.Appid,
		"secret":     r.Config.Appsercet,
	}
	api := "https://api.weixin.qq.com/cgi-bin/stable_token "
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
func (r *Official) SetAccessToken(accessToken string) {
	r.Config.AccessToken = accessToken
}

// Send 服务推送
func (r *Official) Send() {
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
func (r *Official) Handle(callback func(entity entity.OfficialAccountMessageAcceptEntity)) {
	body, err := ioutil.ReadAll(r.Config.Request.Body)
	officialAccountAcceptMessageEntity := entity.OfficialAccountMessageAcceptEntity{}
	if err != nil {
		log.Println(err.Error())
	} else {
		defer r.Config.Request.Body.Close()
		xml.Unmarshal(body, &officialAccountAcceptMessageEntity)
	}
	if officialAccountAcceptMessageEntity.FromUserName != "" {
		callback(officialAccountAcceptMessageEntity)
	}
}

// Qrcode https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
// 	@param	data map[string]interface{}
//	@return	string	二维码地址
func (r *Official) Qrcode(data map[string]interface{}) string {
	var qrcode string
	param, _ := json.Marshal(data)
	api := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + r.Config.AccessToken
	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(param)))
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		var ticket map[string]interface{}
		json.Unmarshal(body, &ticket)
		if ticket["ticket"] != nil {
			qrcode = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(ticket["ticket"].(string))
		}
	}
	return qrcode
}
