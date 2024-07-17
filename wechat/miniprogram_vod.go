package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//miniprogramVod 微短剧
type miniprogramVod struct {
	Config *MiniProgramlConfig
}

func newMiniprogramVod(config *MiniProgramlConfig) miniprogramVod {
	return miniprogramVod{
		Config: config,
	}
}

//媒体实体
type miniprogramMediaEntity struct {
	Errcode   int                        `json:"errcode"`
	Errmsg    string                     `json:"errmsg"`
	MediaInfo miniprogramMediaInfoEntity `json:"media_info"`
}

//媒体列表实体
type miniprogramMediaListEntity struct {
	Errcode       int                        `json:"errcode"`
	Errmsg        string                     `json:"errmsg"`
	MediaInfoList []miniprogramMediaInfoEntity `json:"media_info_list"`
}

//媒体信息
type miniprogramMediaInfoEntity struct {
	MediaId     int64                             `json:"media_id"`
	CreateTime  int64                             `json:"create_time"`
	ExpireTime  int64                             `json:"expire_time"`
	DramaId     int64                             `json:"drama_id"`
	FileSize    string                            `json:"file_size"`
	Duration    int64                             `json:"duration"`
	Name        string                            `json:"name"`
	Description string                            `json:"description"`
	OriginalUrl string                            `json:"original_url"`
	CoverUrl    string                            `json:"cover_url"`
	Mp4Url      string                            `json:"mp4_url"`
	HlsUrl      string                            `json:"hls_url"`
	AuditDetail miniprogramMediaAuditDedailEntity `json:"audit_detail"`
}

//审核信息
type miniprogramMediaAuditDedailEntity struct {
	Status                 int      `json:"status"`
	CreateTime             int64    `json:"create_time"`
	AuditTime              int64    `json:"audit_time"`
	Reason                 string   `json:"reason"`
	EvidenceMaterialIdList []string `json:"evidence_material_id_list"`
}

//GetMediaLink 获取媒体播放链接 https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-3-%E8%8E%B7%E5%8F%96%E5%AA%92%E8%B5%84%E6%92%AD%E6%94%BE%E9%93%BE%E6%8E%A5
//	@param	param	map[string]interface
func (r *miniprogramVod) GetMediaLink(param map[string]interface{}) miniprogramMediaEntity {
	var res miniprogramMediaEntity
	data, _ := json.Marshal(param)
	api := "https://api.weixin.qq.com/wxa/sec/vod/getmedialink?access_token=" + r.Config.AccessToken
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

//GetMedia 获取媒体 https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-2-%E8%8E%B7%E5%8F%96%E5%AA%92%E8%B5%84%E8%AF%A6%E7%BB%86%E4%BF%A1%E6%81%AF
//	@param	param	map[string]interface
func (r *miniprogramVod) GetMedia(param map[string]interface{}) miniprogramMediaEntity {
	var res miniprogramMediaEntity
	data, _ := json.Marshal(param)
	api := "https://api.weixin.qq.com/wxa/sec/vod/getmedia?access_token=" + r.Config.AccessToken
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

//GetMediaList 获取媒体列表 https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-1-%E8%8E%B7%E5%8F%96%E5%AA%92%E8%B5%84%E5%88%97%E8%A1%A8
//	@param	param	map[string]interface
func (r *miniprogramVod) GetMediaList(param map[string]interface{}) miniprogramMediaListEntity {
	var res miniprogramMediaListEntity
	data, _ := json.Marshal(param)
	api := "https://api.weixin.qq.com/wxa/sec/vod/listmedia?access_token=" + r.Config.AccessToken
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
