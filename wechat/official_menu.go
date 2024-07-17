package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type officialMenu struct {
	Config *OfficialConfig
}

func newOfficialMenu(config *OfficialConfig) officialMenu {
	return officialMenu{
		Config: config,
	}
}

// Create 创建菜单 https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html
//	@param param map[string]interface{}	菜单结构Map
//	@return string json字符串
func (r *officialMenu) Create(param map[string]interface{}) map[string]interface{} {
	var res map[string]interface{}
	data, _ := json.Marshal(param)
	api := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + r.Config.AccessToken
	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(data)))
	defer resp.Body.Close()
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
	}
	return res
}

// Get 查询菜单 https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Querying_Custom_Menus.html
//	@return string json字符串
func (r *officialMenu) Get() map[string]interface{} {
	var res map[string]interface{}
	api := "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=" + r.Config.AccessToken
	resp, err := http.Get(api)
	defer resp.Body.Close()
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
	}
	return res
}

// Delete 删除菜单 https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Deleting_Custom-Defined_Menu.html
//	@return string json字符串
func (r *officialMenu) Delete() map[string]interface{} {
	var res map[string]interface{}
	api := "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=" + r.Config.AccessToken
	resp, err := http.Get(api)
	defer resp.Body.Close()
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
	}
	return res
}

// CreateCustom 创建个性化菜单 https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html
//	@param param map[string]interface{}	菜单结构Map
//	@return string json字符串
func (r *officialMenu) CreateCustom(param map[string]interface{}) map[string]interface{} {
	var res map[string]interface{}
	data, _ := json.Marshal(param)
	api := "https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=" + r.Config.AccessToken
	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(data)))
	defer resp.Body.Close()
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
	}
	return res
}

// DeleteCustom 删除个性化菜单 https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html
//	@return string json字符串
func (r *officialMenu) DeleteCustom() map[string]interface{} {
	var res map[string]interface{}
	api := "https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=" + r.Config.AccessToken
	resp, err := http.Get(api)
	defer resp.Body.Close()
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &res)
	}
	return res
}
