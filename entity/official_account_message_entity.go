package entity

import "encoding/xml"

// OfficialAccountMessageEntity 公众号消息发送实体
type OfficialAccountMessageEntity struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	Image        struct {
		MediaId string `xml:"MediaId"`
	}
	Voice struct {
		MediaId string `xml:"MediaId"`
	}
	Video struct {
		MediaId     string `xml:"MediaId"`
		Title       string `xml:"Title"`
		Description string `xml:"Description"`
	}
	Music struct {
		Title        string `xml:"Title"`
		Description  string `xml:"Description"`
		MusicUrl     string `xml:"MusicUrl"`
		HQMusicUrl   string `xml:"HQMusicUrl"`
		ThumbMediaId string `xml:"ThumbMediaId"`
	}
	ArticleCount int64 `xml:"article_count"`
	Articles     []struct {
	}
	XMLName xml.Name `xml:"xml"`
}

type OfficialAccountMessageArticlesItemEntity struct {
	Title       string   `xml:"Title"`
	Description string   `xml:"Description"`
	PicUrl      string   `xml:"PicUrl"`
	Url         string   `xml:"Url"`
	XMLName     xml.Name `xml:"item"`
}
