package wechat

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type officialUser struct {
	Config *OfficialConfig
}

func newOfficialUser(config *OfficialConfig) officialUser {
	return officialUser{
		Config: config,
	}
}

// officialGetUserInfoEntity 用户返回实体
type officialGetUserInfoEntity struct {
	Errcode        int                         `json:"errcode"`
	Errmsg         string                      `json:"errmsg"`
	Subscribe      int                         `json:"subscribe"`
	Openid         string                      `json:"openid"`
	Nickname       string                      `json:"nickname"`
	Sex            int                         `json:"sex"`
	Language       string                      `json:"language"`
	City           string                      `json:"city"`
	Province       string                      `json:"province"`
	Country        string                      `json:"country"`
	Headimgurl     string                      `json:"headimgurl"`
	SubscribeTime  int64                       `json:"subscribe_time"`
	Unionid        string                      `json:"unionid"`
	Remark         string                      `json:"remark"`
	Groupid        int                         `json:"groupid"`
	TagidList      []int                       `json:"tagid_list"`
	SubscribeScene string                      `json:"subscribe_scene"`
	QrScene        int                         `json:"qr_scene"`
	QrSceneStr     string                      `json:"qr_scene_str"`
	UserInfoList   []officialGetUserInfoEntity `json:"user_info_list"`
}

// GetUserInfo 获取用户基本信息
//	@param	openId	string	用户openid
func (r *officialUser) GetUserInfo(openId string) officialGetUserInfoEntity {
	var res officialGetUserInfoEntity
	api := "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + r.Config.AccessToken + "&openid=" + openId + "&lang=zh_CN"
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

// officialGetUserListEntity 用户返回实体
type officialGetUserListEntity struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Total      int    `json:"total"`
	Count      string `json:"count"`
	NextOpenid string `json:"next_openid"`
	Data       struct {
		OpenId []string
	} `json:"data"`
}

// GetUserList 获取用户列表
//	@param	nextOpenId	string	第一个拉取的OPENID，不填默认从头开始拉取
func (r *officialUser) GetUserList(nextOpenId string) officialGetUserListEntity {
	var res officialGetUserListEntity
	api := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + r.Config.AccessToken + "&next_openid=" + nextOpenId
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

// Oauth 微信授权 https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
//	@param	redirect	string	重定向地址
//	@param	scope	string	授权类型	snsapi_base和snsapi_userinfo
//	@param	state	string	授权类型	重定向之后会带上state参数
func (r *officialUser) Oauth(redirect string, scope string, state string) {
	api := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + r.Config.Appid + "&redirect_uri=" + url.QueryEscape(redirect) + "&response_type=code&scope=" + scope + "&state=" + state + "#wechat_redirect"
	http.Redirect(r.Config.ResponseWrite, r.Config.Request, api, http.StatusFound)
}

//获取用户授权token实体
type officialGetUserTokenEntity struct {
	Erroce       int64  `json:"erroce"`
	ErrMsg       string `json:"err_msg"`
	AccessToken  string `json:"access_token"`
	ExpireIn     string `json:"expire_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string
}

// GetUserToken 获取用户
//	@param	code	string
func (r *officialUser) GetUserToken(code string) officialGetUserTokenEntity {
	var res officialGetUserTokenEntity
	api := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + r.Config.Appid + "&secret=" + r.Config.Appsercet + "&code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(api)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
	} else {
		read, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(read, &res)
	}
	return res
}

//公众号授权获取用户信息实体
type officialGetUserEntity struct {
	Errcode    int64    `json:"errcode"`
	Errmsg     string   `json:"errmsg"`
	Openid     string   `json:"openid"`
	NickName   string   `json:"nick_name"`
	Sex        int64    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

// GetUser 获取用户
//	@param	token	string
//	@param	openid	string
func (r *officialUser) GetUser(token string, openid string) officialGetUserEntity {
	var res officialGetUserEntity
	api := "https://api.weixin.qq.com/sns/userinfo?access_token=" + token + "&openid=" + openid + "&lang=zh_CN"
	resp, err := http.Get(api)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
	} else {
		read, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(read, &res)
	}
	return res
}
