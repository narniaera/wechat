package entity

// MiniprogramMessageAcceptEntity 小程序消息接收实体
type MiniprogramMessageAcceptEntity struct {
	Errcode       int    `json:"errcode"`
	Errmsg        string `json:"errmsg"`
	ToUserName    string
	FromUserName  string
	CreateTime    int64
	MsgType       string
	Event         string
	Appid         string `json:"appid"`
	TraceId       string `json:"trace_id"`
	Version       int64  `json:"version"`
	OpenId        string
	OutTradeNo    string
	Env           int
	WeChatPayInfo miniprogramWeChatPayVirtualInfo
	CoinInfo      miniprogramWeChatPayVirtualCoinInfo
	Detail        []miniprogramMessageDetailEntity
	Result        miniprogramMessageResultEntity `json:"result"`
}

type miniprogramMessageResultEntity struct {
	Suggest string `json:"suggest"`
	Label   int64  `json:"label"`
}
type miniprogramMessageDetailEntity struct {
	Strategy string `json:"strategy"`
	Errcode  int64  `json:"errcode"`
	Suggest  string `json:"suggest"`
	Label    int64  `json:"label"`
	Prob     int64  `json:"prob"`
}

//虚拟支付订单信息
type miniprogramWeChatPayVirtualInfo struct {
	MchOrderNo    string
	TransactionId string
}

//虚拟支付订单代币参数信息
type miniprogramWeChatPayVirtualCoinInfo struct {
	Quantity    int
	OrigPrice   int
	ActualPrice int
	Attach      string
}
