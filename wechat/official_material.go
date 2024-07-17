package wechat

import (
	"bytes"
	"encoding/json"
	"github.com/narniaera/wechat/entity"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type officialMaterial struct {
	Config *OfficialConfig
}

func newOfficialMaterial(config *OfficialConfig) officialMaterial {
	return officialMaterial{
		Config: config,
	}
}

type materialUploadEntity struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
	Url       string `json:"url"`
}

type materialEntity struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	MediaId string `json:"media_id"`
}

// AddNews 新增永久素材
func (r *officialMaterial) AddNews(param map[string]interface{}) materialEntity {
	data, _ := json.Marshal(param)
	api := "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=" + r.Config.AccessToken
	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(data)))
	defer resp.Body.Close()
	var res materialEntity
	if err != nil {
		log.Println(err.Error())
	} else {
		read, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(read, &res)
	}
	return res
}

// UploadTemp 上传临时素材 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
//	@param	fileType	string	文件类型
//	@param	file	*multipart.FileHeader	表单文件流
func (r *officialMaterial) UploadTemp(fileType string, file *multipart.FileHeader) materialUploadEntity {
	//打开文件
	fh, err := os.Open(file.Filename)
	if err != nil {
		log.Println(err.Error())
		return materialUploadEntity{}
	}
	defer fh.Close()
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("media", filepath.Base(file.Filename))
	if err != nil {
		return materialUploadEntity{}
	}
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		log.Println(err.Error())
		return materialUploadEntity{}
	}
	bodyWriter.Close()
	//upload
	api := "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=" + r.Config.AccessToken + "&type=" + fileType
	req, err := http.NewRequest("POST", api, bodyBuf)
	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())
	urlQuery := req.URL.Query()
	if err != nil {
		log.Println(err.Error())
		return materialUploadEntity{}
	}
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return materialUploadEntity{}
	}
	defer resp.Body.Close()
	read, _ := ioutil.ReadAll(resp.Body)
	res := materialUploadEntity{}
	json.Unmarshal(read, &res)
	return res
}

// Upload 上传永久素材	https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
//	@param	fileType	string	文件类型
//	@param	file	*multipart.FileHeader	表单文件流
func (r *officialMaterial) Upload(fileType string, file *multipart.FileHeader, param map[string]string) materialUploadEntity {
	//打开文件
	fh, err := file.Open()
	if err != nil {
		log.Println(err.Error())
		return materialUploadEntity{}
	}
	defer fh.Close()
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("media", filepath.Base(file.Filename))
	if err != nil {
		return materialUploadEntity{}
	}
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		log.Println(err.Error())
		return materialUploadEntity{}
	}
	//循环字段
	if len(param) != 0 {
		for k, v := range param {
			bodyWriter.WriteField(k, v)
		}
	}

	bodyWriter.Close()
	//upload
	api := "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=" + r.Config.AccessToken + "&type=" + fileType
	req, err := http.NewRequest("POST", api, bodyBuf)
	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())
	urlQuery := req.URL.Query()
	if err != nil {
		log.Println(err.Error())
		return materialUploadEntity{}
	}
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return materialUploadEntity{}
	}
	defer resp.Body.Close()
	read, _ := ioutil.ReadAll(resp.Body)
	res := materialUploadEntity{}
	json.Unmarshal(read, &res)
	return res
}

// UploadImg 上传图文消息内的图片获取URL
func (r *officialMaterial) UploadImg(file *multipart.FileHeader) materialUploadEntity {
	//打开文件
	fh, err := os.Open(file.Filename)
	if err != nil {
		log.Println(err.Error())
		return materialUploadEntity{}
	}
	defer fh.Close()
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("media", filepath.Base(file.Filename))
	if err != nil {
		return materialUploadEntity{}
	}
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		log.Println(err.Error())
		return materialUploadEntity{}
	}
	bodyWriter.Close()
	//upload
	api := "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=" + r.Config.AccessToken
	req, err := http.NewRequest("POST", api, bodyBuf)
	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())
	urlQuery := req.URL.Query()
	if err != nil {
		log.Println(err.Error())
		return materialUploadEntity{}
	}
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return materialUploadEntity{}
	}
	defer resp.Body.Close()
	read, _ := ioutil.ReadAll(resp.Body)
	res := materialUploadEntity{}
	json.Unmarshal(read, &res)
	return res
}

// 素材获取实体
type materialGetEntity struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	MediaId  string `json:"media_id"`
	VideoUrl int64  `json:"video_url"`
}

// GetTempMedia 获取临时素材
//	@param mediaId 媒体id
//	@return materialGetEntity
func (r *officialMaterial) GetTempMedia(mediaId string) materialGetEntity {
	api := "https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=" + r.Config.AccessToken + "&media_id=" + mediaId
	resp, err := http.Get(api)
	defer resp.Body.Close()
	var res materialGetEntity
	if err != nil {
		log.Println(err.Error())
	} else {
		read, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(read, &res)
	}
	return res
}

//获取永久素材实体
type getMaterialEntity struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	NewsItem []struct {
		Title            string `json:"title"`              //图文消息的标题
		ThumbMediaId     string `json:"thumb_media_id"`     //图文消息的封面图片素材id（必须是永久mediaID）
		ShowCoverPic     int    `json:"show_cover_pic"`     //是否显示封面，0为false，即不显示，1为true，即显示
		Author           string `json:"author"`             //作者
		Digest           string `json:"digest"`             //图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空
		Content          string `json:"content"`            //图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS
		Url              string `json:"url"`                //图文页的URL
		ContentSourceUrl string `json:"content_source_url"` //图文消息的原文地址，即点击“阅读原文”后的URL
	} `json:"news_item"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DownUrl     string `json:"down_url"`
}

// GetMedia 获取永久素材 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (r *officialMaterial) GetMedia(mediaId string) getMaterialEntity {
	data := map[string]string{
		"media_id": mediaId,
	}
	param, _ := json.Marshal(data)
	api := "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=" + r.Config.AccessToken
	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(param)))
	defer resp.Body.Close()
	var res getMaterialEntity
	if err != nil {
		log.Println(err.Error())
	} else {
		read, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(read, &res)
	}
	return res
}

// Delete 删除永久素材 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Deleting_Permanent_Assets.html
func (r *officialMaterial) Delete(mediaId string) entity.ResultEntity {
	data := map[string]string{
		"media_id": mediaId,
	}
	param, _ := json.Marshal(data)
	api := "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=" + r.Config.AccessToken
	resp, err := http.Post(api, "application/json", bytes.NewBufferString(string(param)))
	defer resp.Body.Close()
	var res entity.ResultEntity
	if err != nil {
		log.Println(err.Error())
	} else {
		read, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(read, &res)
	}
	return res
}
