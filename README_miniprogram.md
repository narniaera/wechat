# g公众号

 https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html

## 1. 初始化

- 初始化

  ```go
  miniprogram := wechat.NewMiniProgram(&wechat.MiniProgramlConfig{
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
  miniprogram.GetAccessToken(true)
  //设置access_token 用于缓存数据读取在设值
  miniprogram.SetAccessToken("")
  ```

## 2. 用户登录

- code2Session

  ```go
  miniprogram.User.Jscode2session("code码")
  ```
