# g公众号

 https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html

## 1. 初始化

- 初始化

  ```go
  official := wechat.NewOfficial(&wechat.OfficialConfig{
      Appid:     "",
      Appsercet: "",
      Token:     "",
      Request:   c.Request,
      ResponseWrite: c.Writer,
  })
  ```

- 获取、设置access_token

  ```go
  //获取access_token
  official.GetAccessToken(true)
  //设置access_token 用于缓存数据读取在设值
  official.SetAccessToken("")
  ```

## 2. 消息推送

- token验证

  ```go
  official.Send()
  ```

- 消息推送

  ```go
  official.Handle(func(message entity.OfficialAccountMessageAcceptEntity) {
      //被动消息实体
      officialAccountMessageEntity := entity.OfficialAccountMessageEntity{}
      officialAccountMessageEntity.ToUserName = message.FromUserName
      officialAccountMessageEntity.FromUserName = message.ToUserName
      //事件s判断
      if message.MsgType == "event" {
          if message.Event == "subscribe" {
              fmt.Println("关注")
              if message.EventKey != "" {
                  fmt.Println("二维码关注")
              }
          } else if message.Event == "unscribe" {
              fmt.Println("取消关注")
          }
      } else if message.MsgType == "text" {
          officialAccountMessageEntity.MsgType = "text"
          officialAccountMessageEntity.Content = "aaaaa"
          official.Message.Push(officialAccountMessageEntity)
      } else if message.MsgType == "image" {
      	fmt.Println("图片")
      } else if message.MsgType == "voice" {
      	fmt.Println("语音")
      } else if message.MsgType == "video" {
      	fmt.Println("视频")
      } else if message.MsgType == "shortvideo" {
      	fmt.Println("小视频")
      } else if message.MsgType == "link" {
      	fmt.Println("链接")
      }
      c.String(http.StatusOK, "success")
      return
  })
  official.Send()
  ```

- 回复客服消息

  ```go
  	data := map[string]interface{}{
  		"touser":  "oPllS6MKsjEpvVEknQJgPiWZTZk0",
  		"msgtype": "news",
  		"news": map[string]interface{}{
  			"articles": []map[string]string{
  				{
  					"title":       "新闻标题",
  					"description": "新闻描述",
  					"url":         "https://baidu.com",
  					"picurl":      "https://www.baidu.com/img/PCfb_5bf082d29588c07f842ccde3f97243ea.png",
  				},
  			},
  		},
  	}
  	r := wechat.Message.Custom(data)
  ```

  

## 3. 自定义菜单

- 创建菜单 

  ```go
  buttons := map[string]interface{}{
      "button": []map[string]interface{}{
          {
              "type": "view",
              "name": "BaiDu",
              "url":  "https://baidu.com",
              "sub_button": []map[string]interface{}{
                  {
                      "type": "view",
                      "name": "搜索",
                      "url":  "http://www.soso.com/"
                  },
              },
          },
      },
  }
  official.Menu.Create(buttons)
  ```

- 查询菜单

  ```go
  official.Menu.Get()
  ```

- 删除菜单

  ```go
  official.Menu.Get()
  ```

- 创建个性化菜单

  ```go
  official.Menu.CreateCustom(buttons)
  ```

- 删除个性化菜单

  ```go
  official.MenuDeleteCustom()
  ```

## 4. 模板消息

- 获取模板列表

  ```go
  official.Template.GetAllTemplate()
  ```

- 删除模板

  ```go
  data := map[string]string{
  	"template_id": "template_id",
  }
  res := official.Template.Delete(data)
  ```

- 发送模板消息

  ```go
  data := map[string]interface{}{
      "touser":      "oPllS6MKsjEpvVEknQJgPiWZTZk0",
      "template_id": "GLLgsh2tAicyAP03cLbsGOvYHwlbIDjrD7JDYuKKLG4",
      "data": map[string]interface{}{
      	"title": map[string]string{
      		"value": "这个是标题",
      	},
      	"desc": map[string]string{
      		"value": "这个是描述",
      	},
      },
  }
  official.Template.Send(data)
  ```

## 5. 用户管理

- 获取用户基本信息

  ```go
  official.User.GetUserInfo("openid")
  ```

- 获取用户列表

  ```go
  official.User.GetUserList("next_open_id")
  ```

## 6. 创建二维码

- ```go
  //https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
  data := map[string]interface{}{
      "expire_seconds": 3600,
      "action_name":    "QR_SCENE",
      "action_info": map[string]interface{}{
          "scene": map[string]interface{}{
          "test": 123,
          },
      },
  }
  official.Qrcode(data)
  ```

## 7. 网页授权

- ```go
  	official := wechat.NewOfficial(&wechat.OfficialConfig{
  		Appid:         "wxc8399122cea914bc",
  		Appsercet:     "ab47cc15caaffb9fcc7326113f66d629",
  		Token:         "token",
  		Request:       c.Request,
  		ResponseWrite: c.Writer,
  	})
  	if c.Query("code") != "" {
  		oauth := official.User.GetUserToken(c.Query("code"))
  		user := official.User.GetUser(oauth["access_token"].(string), oauth["openid"].(string))
  		fmt.Println(user)
  	} else {
  		official.User.Oauth("http://au5hfa.natappfree.cc/login/wechatOauth", "snsapi_userinfo", "")
  	}
  ```

## 8. 素材管理

- 上传临时素材

  ```go
  f, _ := c.FormFile("media")
  official.Material.Upload("image",f)
  ```

- 上传永久素材

  ```go
  f, _ := c.FormFile("media")
  official.Material.Upload("image", f, map[string]string{"description": ""})
  ```

- 上传图片素材

  ```go
  f, _ := c.FormFile("media")
  official.Material.UploadImg("image", f, map[string]string{"description": ""})
  ```

- 删除永久素材

  ```go
  official.Material.Delete("media_id")
  ```

  