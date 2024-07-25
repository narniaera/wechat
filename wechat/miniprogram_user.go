package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type miniprogramUser struct {
	Config *MiniProgramlConfig
}

func newMiniprogramUser(config *MiniProgramlConfig) miniprogramUser {
	return miniprogramUser{
		Config: config,
	}
}

// miniprogramUserEntity 用户code2Session返回实体 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
type miniprogramUserEntity struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	SessionKey string `json:"session_key"`
}

// miniprogramUserPhoneNumberEntity 用户getPhoneNumber返回实体 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
type miniprogramUserPhoneNumberEntity struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	PhoneInfo struct {
		PhoneNumber     string `json:"phoneNumber"`
		PurePhoneNumber string `json:"purePhoneNumber"`
		CountryCode     int    `json:"countryCode"`
		Watermark       struct {
			Timestamp int64  `json:"timestamp"`
			Appid     string `json:"appid"`
		} `json:"watermark"`
	} `json:"phone_info"`
}

// GetUserInfo 获取用户基本信息
//	@param	code	string	小程序code码
func (r *miniprogramUser) Jscode2session(code string) miniprogramUserEntity {
	var res miniprogramUserEntity
	api := "https://api.weixin.qq.com/sns/jscode2session?appid=" + r.Config.Appid + "&secret=" + r.Config.Appsercet + "&js_code=" + code + "&grant_type=authorization_code"
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

// GetPhoneNumber 获取用户手机号
//	@param	code	string	小程序code码
func (r *miniprogramUser) GetPhoneNumber(code string) miniprogramUserPhoneNumberEntity {
	var res miniprogramUserPhoneNumberEntity
	params := map[string]interface{}{
		"code": code,
	}
	data, _ := json.Marshal(params)
	api := "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=" + r.Config.AccessToken
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
