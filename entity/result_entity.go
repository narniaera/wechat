package entity

// ResultEntity 返回实体
type ResultEntity struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ResultAccessTokenEntity 返回实体
type ResultAccessTokenEntity struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
