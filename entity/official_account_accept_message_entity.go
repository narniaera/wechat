package entity

// OfficialAccountMessageAcceptEntity 公众号消息接收实体
type OfficialAccountMessageAcceptEntity struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	Ticket       string `xml:"Ticket"`
	MsgID string `xml:"msg_id"`
	Status string `xml:"status"`
}
