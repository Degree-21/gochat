# 微信公众号（Official Accounts）

```go
"github.com/shenghui0779/gochat"
"github.com/shenghui0779/gochat/oa"
```

### 初始化公众号实例

```go
wxoa := gochat.NewOA(appid, appsecret)

// 如果启用了服务器配置，需要设置配置项
wxoa.SetServerConfig(token, encodingAESKey)

// 如果需要消息回复，需要设置原始ID（开发者微信号）
wxoa.SetOriginID(originID)
```

### 网页授权

```go
// 生成网页授权URL（请使用 URLEncode 对 redirectURL 进行处理）
wxoa.AuthURL(scope, redirect_url)

// 获取网页授权AccessToken
wxoa.Code2AuthToken(ctx, code)

// 刷新网页授权AccessToken
wxoa.RefreshAuthToken(ctx, refresh_token)

// 检验授权凭证（access_token）是否有效
wxoa.Do(ctx, access_token, oa.CheckAuthToken(openid))

// 获取授权用户信息（注意：使用网页授权的access_token）
wxoa.Do(ctx, access_token, oa.GetAuthUser(dest, openid))

// 获取普通AccessToken
wxoa.AccessToken(ctx)
```

### 自定义菜单

```go
// 创建自定义菜单
wxoa.Do(ctx, access_token, oa.CreateMenu(buttons...))

// 创建个性化菜单
wxoa.Do(ctx, access_token, oa.CreateConditionalMenu(match_rule, buttons...))

// 查询自定义菜单
wxoa.Do(ctx, access_token, oa.GetMenu(dest))

// 删除自定义菜单
wxoa.Do(ctx, access_token, oa.DeleteMenu())

// 删除个性化菜单
wxoa.Do(ctx, access_token, oa.DeleteConditionalMenu(menu_id))
```

### 消息

```go
// 获取模板列表
wxoa.Do(ctx, access_token, oa.GetTemplateList(dest))

// 删除模板
wxoa.Do(ctx, access_token, oa.DeleteTemplate(template_id))

// 发送模板消息
wxoa.Do(ctx, access_token, oa.SendTemplateMessage(openid, msg))

// 发送订阅消息
wxoa.Do(ctx, access_token, oa.SendSubscribeMessage(openid, scene, title, msg))
```

### 用户管理

```go
// 获取关注用户信息
wxoa.Do(ctx, access_token, oa.GetSubscriberInfo(dest, openid))

// 批量关注用户信息
wxoa.Do(ctx, access_token, oa.BatchGetSubscribers(dest, openids...)

// 获取关注用户列表
wxoa.Do(ctx, access_token, oa.GetSubscriberList(dest, next_openid)

// 获取用户黑名单列表
wxoa.Do(ctx, access_token, oa.GetBlackList(dest, begin_openid)

// 拉黑用户
wxoa.Do(ctx, access_token, oa.BlackSubscribers(openids...))

// 取消拉黑用户
wxoa.Do(ctx, access_token, oa.UnBlackSubscribers(openids...))

// 设置用户备注名（该接口暂时开放给微信认证的服务号）
wxoa.Do(ctx, access_token, oa.SetUserRemark(openid, remark))
```

### 二维码

```go
// 创建临时二维码（expireSeconds：二维码有效时间，最大不超过2592000秒（即30天），不填，则默认有效期为30秒。）
wxoa.Do(ctx, access_token, oa.CreateTempQRCode(dest, sence_id, expire_seconds...))

// 创建永久二维码（expireSeconds：二维码有效时间，最大不超过2592000秒（即30天），不填，则默认有效期为30秒。）
wxoa.Do(ctx, access_token, oa.CreatePermQRCode(dest, sence_id, expire_seconds...))
```

### 素材管理

```go
// 上传临时素材
wxoa.Do(ctx, access_token, oa.UploadMedia(dest, media_type, filename))

// 新增永久图文素材（公众号的素材库保存总数量有上限：图文消息素材、图片素材上限为100000，其他类型为1000）
wxoa.Do(ctx, access_token, oa.AddNews(dest, articles...))

// 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制，图片仅支持jpg/png格式，大小必须在1MB以下）
wxoa.Do(ctx, access_token, oa.UploadNewsImage(dest, filename))

// 新增其他类型永久素材（支持图片、音频、缩略图）
wxoa.Do(ctx, access_token, oa.AddMaterial(dest, media_type, filename))

// 上传视频永久素材
wxoa.Do(ctx, access_token, oa.UploadVideo(dest, filename, title, introduction))

// 删除永久素材
wxoa.Do(ctx, access_token, oa.DeleteMaterial(media_id))
```

### 图像处理

```go
// 图片智能裁切
wxoa.Do(ctx, access_token, oa.AICrop(dest, filename))
wxoa.Do(ctx, access_token, oa.AICropByURL(dest, imgURL))

// 条码/二维码识别
wxoa.Do(ctx, access_token, oa.ScanQRCode(dest, filename))
wxoa.Do(ctx, access_token, oa.ScanQRCodeByURL(dest, imgURL))

// 图片高清化
wxoa.Do(ctx, access_token, oa.SuperreSolution(dest, filename))
wxoa.Do(ctx, access_token, oa.SuperreSolutionByURL(dest, imgURL))
```

### OCR

```go
// 银行卡识别
wxoa.Do(ctx, access_token, oa.OCRBankCard(dest, mode, filename))
wxoa.Do(ctx, access_token, oa.OCRBankCardByURL(dest, mode, imgURL))

// 营业执照识别
wxoa.Do(ctx, access_token, oa.OCRBusinessLicense(dest, mode, filename))
wxoa.Do(ctx, access_token, oa.OCRBusinessLicenseByURL(dest, mode, imgURL))

// 身份证前面识别
wxoa.Do(ctx, access_token, oa.OCRIDCardFront(dest, mode, filename))
wxoa.Do(ctx, access_token, oa.OCRIDCardFrontByURL(dest, mode, imgURL))

// 身份证背面识别
wxoa.Do(ctx, access_token, oa.OCRIDCardBack(dest, mode, filename))
wxoa.Do(ctx, access_token, oa.OCRIDCardBackByURL(dest, mode, imgURL))

// 通用印刷体识别
wxoa.Do(ctx, access_token, oa.OCRPrintedText(dest, mode, filename))
wxoa.Do(ctx, access_token, oa.OCRPrintedTextByURL(dest, mode, imgURL))

// 行驶证识别
wxoa.Do(ctx, access_token, oa.OCRVehicleLicense(dest, mode, filename))
wxoa.Do(ctx, access_token, oa.OCRVehicleLicenseByURL(dest, mode, imgURL))

// 车牌号识别
wxoa.Do(ctx, access_token, oa.OCRPlateNumber(dest, mode, filename))
wxoa.Do(ctx, access_token, oa.OCRPlateNumberByURL(dest, mode, imgURL))
```

### JSSDK

```go
// GetJSSDKTicket 获取 JS-SDK ticket (注意：使用普通access_token)
wxoa.Do(ctx, access_token, oa.GetJSSDKTicket(dest, ticket_type))

// 生成 JS-SDK 签名
wxoa.JSSDKSign(jsapi_ticket, url)
```

### 消息事件

```go
// 验证消息事件签名
wxoa.VerifyEventSign(signature, items...)

// 事件消息解密
wxoa.DecryptEventMessage(msg_encrypt)
```

### 消息回复

```go
// 回复文本消息
wxoa.Reply(openid, oa.NewTextReply(content))

// 回复图片消息
wxoa.Reply(openid, oa.NewImageReply(media_id))

// 回复语音消息
wxoa.Reply(openid, oa.NewVoiceReply(media_id))

// 回复视频消息
wxoa.Reply(openid, oa.NewVideoReply(media_id, title, desc))

// 回复音乐消息
wxoa.Reply(openid, oa.NewMusicReply(media_id, title, desc, music_url, HQ_music_url))

// 回复图文消息
wxoa.Reply(openid, oa.NewNewsReply(count, articles...))

// 回复客服消息
wxoa.Reply(openid, oa.NewTransfer2KFReply(kf_account...))
```