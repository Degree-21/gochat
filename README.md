# gochat

[![golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![GitHub release](https://img.shields.io/github/release/shenghui0779/gochat.svg)](https://github.com/shenghui0779/gochat/releases/latest)
[![pkg.go.dev](https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/shenghui0779/gochat)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)


## 该项目进入维护阶段，不再新增功能
## 按功能模块拆分成一下三个独立的新SDK

- [微信支付](https://github.com/shenghui0779/wechat_pay)
- [微信小程序](https://github.com/shenghui0779/wechat_mp)
- [微信公众号](https://github.com/shenghui0779/wechat_oa)

微信 SDK for Go

| 目录 | 对应         | 功能                                               |
| ---- | ------------ | -------------------------------------------------- |
| /mch | 微信商户平台 | 下单、支付、退款、查询、委托代扣、企业付款、企业红包 等 |
| /pub | 微信公众平台 | 网页授权、菜单、模板消息、消息回复、用户管理、消息转客服 等 |
| /mp  | 微信小程序   | 小程序授权、用户数据解析、消息发送、二维码生成、消息 等 |

## 获取

```sh
go get github.com/shenghui0779/gochat
```

## 文档

- [API Reference](https://pkg.go.dev/github.com/shenghui0779/gochat)
- [Wiki](https://github.com/shenghui0779/gochat/wiki)

## 说明

- 支持 Go1.11+
- 注意：因 `access_token` 每日获取次数有限且含有效期，故服务端应妥善保存 `access_token` 并定时刷新
- 配合 [yiigo](https://github.com/shenghui0779/yiigo) 使用，可以更方便的操作 `MySQL`、`MongoDB` 与 `Redis` 等

**Enjoy 😊**

