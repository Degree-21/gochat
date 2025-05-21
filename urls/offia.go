package urls

// oauth
const (
	Oauth2Authorize  = "https://open.weixin.qq.com/connect/oauth2/authorize"
	SubscribeMsgAuth = "https://mp.weixin.qq.com/mp/subscribemsg"
)

// cgi-bin
const (
	OffiaCgiBinAccessToken = "https://api.weixin.qq.com/cgi-bin/token"
	OffiaCgiBinTicket      = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
)

// menu
const (
	OffiaMenuCreate            = "https://api.weixin.qq.com/cgi-bin/menu/create"
	OffiaGetCurSelfMenuInfo    = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info"
	OffiaMenuAddConditional    = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"
	OffiaMenuTryMatch          = "https://api.weixin.qq.com/cgi-bin/menu/trymatch"
	OffiaMenuGet               = "https://api.weixin.qq.com/cgi-bin/menu/get"
	OffiaMenuDelete            = "https://api.weixin.qq.com/cgi-bin/menu/delete"
	OffiaMenuDeleteConditional = "https://api.weixin.qq.com/cgi-bin/menu/delconditional"
)

// sns
const (
	OffiaSnsCode2Token         = "https://api.weixin.qq.com/sns/oauth2/access_token"
	OffiaSnsCheckAccessToken   = "https://api.weixin.qq.com/sns/auth"
	OffiaSnsRefreshAccessToken = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	OffiaSnsUserInfo           = "https://api.weixin.qq.com/sns/userinfo"
)

// subscriber
const (
	OffiaUserGet      = "https://api.weixin.qq.com/cgi-bin/user/info"
	OffiaUserBatchGet = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
	OffiaUserList     = "https://api.weixin.qq.com/cgi-bin/user/get"
	// 公众号标签
	OffiaTagCreate        = "https://api.weixin.qq.com/cgi-bin/tags/create"
	OffiaTagGet           = "https://api.weixin.qq.com/cgi-bin/tags/get"
	OffiaTagBatchTagging  = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging"
	OffiaBlackListGet     = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist"
	OffiaBatchBlackList   = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist"
	OffiaBatchUnBlackList = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist"
	OffiaUserRemarkSet    = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark"
)

// message
const (
	OffiaSetIndustry              = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry"
	OffiaGetIndustry              = "https://api.weixin.qq.com/cgi-bin/template/get_industry"
	OffiaTemplateAdd              = "https://api.weixin.qq.com/cgi-bin/template/api_add_template"
	OffiaGetAllPrivateTemplate    = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template"
	OffiaDelPrivateTemplate       = "https://api.weixin.qq.com/cgi-bin/template/del_private_template"
	OffiaTemplateMsgSend          = "https://api.weixin.qq.com/cgi-bin/message/template/send"
	OffiaSubscribeTemplateMsgSend = "https://api.weixin.qq.com/cgi-bin/message/template/subscribe"
)

// popularize
const (
	OffiaQRCodeCreate     = "https://api.weixin.qq.com/cgi-bin/qrcode/create"
	OffiaQRCodeShow       = "https://mp.weixin.qq.com/cgi-bin/showqrcode"
	OffiaShortURLGenerate = "https://api.weixin.qq.com/cgi-bin/shorturl"
)

// media
const (
	OffiaMediaUpload     = "https://api.weixin.qq.com/cgi-bin/media/upload"
	OffiaMediaUploadNews = "https://api.weixin.qq.com/cgi-bin/media/uploadnews"
	OffiaMediaGet        = "https://api.weixin.qq.com/cgi-bin/media/get"
	OffiaNewsAdd         = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	OffiaNewUpdate       = "https://api.weixin.qq.com/cgi-bin/material/update_news"
	OffiaNewsImageUpload = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
	OffiaMaterialAdd     = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	OffiaMaterialDelete  = "https://api.weixin.qq.com/cgi-bin/material/del_material"
	OffiaMaterialGet     = "https://api.weixin.qq.com/cgi-bin/material/get_material"
)

// image
const (
	OffiaAICrop          = "https://api.weixin.qq.com/cv/img/aicrop"
	OffiaScanQRCode      = "https://api.weixin.qq.com/cv/img/qrcode"
	OffiaSuperreSolution = "https://api.weixin.qq.com/cv/img/superresolution"
)

// ocr
const (
	OffiaOCRIDCard          = "https://api.weixin.qq.com/cv/ocr/idcard"
	OffiaOCRBankCard        = "https://api.weixin.qq.com/cv/ocr/bankcard"
	OffiaOCRPlateNumber     = "https://api.weixin.qq.com/cv/ocr/platenum"
	OffiaOCRDriverLicense   = "https://api.weixin.qq.com/cv/ocr/drivinglicense"
	OffiaOCRVehicleLicense  = "https://api.weixin.qq.com/cv/ocr/driving"
	OffiaOCRBusinessLicense = "https://api.weixin.qq.com/cv/ocr/bizlicense"
	OffiaOCRComm            = "https://api.weixin.qq.com/cv/ocr/comm"
)

// KF
const (
	OffiaKFAccountList   = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist"
	OffiaKFOnlineList    = "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist"
	OffiaKFAccountAdd    = "https://api.weixin.qq.com/customservice/kfaccount/add"
	OffiaKFInvite        = "https://api.weixin.qq.com/customservice/kfaccount/inviteworker"
	OffiaKFAccountUpdate = "https://api.weixin.qq.com/customservice/kfaccount/update"
	OffiaKFAvatarUpload  = "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg"
	OffiaKFDelete        = "https://api.weixin.qq.com/customservice/kfaccount/del"
	OffiaKFSessionCreate = "https://api.weixin.qq.com/customservice/kfsession/create"
	OffiaKFSessionClose  = "https://api.weixin.qq.com/customservice/kfsession/close"
	OffiaKFSessionGet    = "https://api.weixin.qq.com/customservice/kfsession/getsession"
	OffiaKFSessionList   = "https://api.weixin.qq.com/customservice/kfsession/getsessionlist"
	OffiaKFWaitCase      = "https://api.weixin.qq.com/customservice/kfsession/getwaitcase"
	OffiaKFMsgRecordList = "https://api.weixin.qq.com/customservice/msgrecord/getmsglist"
	OffiaKFMsgSend       = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	OffiaSetTyping       = "https://api.weixin.qq.com/cgi-bin/message/custom/typing"
)

// card
const (
	// 卡券创建接口
	CardCreate = "https://api.weixin.qq.com/card/create"
	// 查询 卡券详情
	CardGet = "https://api.weixin.qq.com/card/get"
	// 删除 卡券
	CardDelete = "https://api.weixin.qq.com/card/delete"
	// 修改库存接口
	CardModifyStock = "https://api.weixin.qq.com/card/modifystock"
	// 卡券二维码ticker接口
	CardQrcodeCreate = "https://api.weixin.qq.com/card/qrcode/create"
	// 设置测试卡券白名单接口
	CardWhitelist = "https://api.weixin.qq.com/card/testwhitelist/set"
	// 核销卡券接口
	CardCodeConsume = "https://api.weixin.qq.com/card/code/consume"
	// 核销卡券 查询 code 接口
	CardCodeGet = "https://api.weixin.qq.com/card/code/get"
	// 修改卡券code
	CardCodeUpdate = "https://api.weixin.qq.com/card/code/update"
	// 获取用户已领取的卡券
	CardUserList = "https://api.weixin.qq.com/card/user/getcardlist"
	// 批量获取卡券列表
	CardBatchGet = "https://api.weixin.qq.com/card/batchget"
	// 更改卡券信息
	CardUpdate = "https://api.weixin.qq.com/card/update"
	// 设置卡券失效接口
	CardCodeUnavailable = "https://api.weixin.qq.com/card/code/unavailable"
	// 拉取卡券概括数据
	CardBizUinInfo = "https://api.weixin.qq.com/datacube/getcardbizuininfo"
	// code解码接口
	CardCodeDecrypt = "https://api.weixin.qq.com/card/code/decrypt"
	//获取免费券数据接口（优惠券、团购券、折扣券、礼品券）
	CardFreeGet = "https://api.weixin.qq.com/datacube/getcardcardinfo"
	//创建卡券货架
	CardLandingPageCreate = "https://api.weixin.qq.com/card/landingpage/create"
)

const (
	//激活会员卡
	VipCardActivate = "https://api.weixin.qq.com/card/membercard/activate"
	//获取会员卡信息
	GetVipCardInfo = "https://api.weixin.qq.com/card/membercard/userinfo/get"
	//更新会员用户信息
	UpdateVipCardUserInfo = "https://api.weixin.qq.com/card/membercard/updateuser"
	//更新会员卡信息
	UpdateVipCardInfo = "https://api.weixin.qq.com/card/update"
)

// publish 发布能力
//https://developers.weixin.qq.com/doc/offiaccount/Publish/Publish.html
const (
	// 发布
	FreePublishSubmit = "https://api.weixin.qq.com/cgi-bin/freepublish/submit"
	// 轮训获取结果
	FreePublishGet = "https://api.weixin.qq.com/cgi-bin/freepublish/get"
	// 删除
	FreePublishDelete = "https://api.weixin.qq.com/cgi-bin/freepublish/delete"
	//通过 article_id 获取已发布文章
	FreePublishGetArticle = "https://api.weixin.qq.com/cgi-bin/freepublish/getarticle"
	//获取成功发布列表
	FreePublishBatchGet = "https://api.weixin.qq.com/cgi-bin/freepublish/batchget"
)

//draft https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html
// 草稿箱能力
const (
	// 添加草稿
	DraftAdd = "https://api.weixin.qq.com/cgi-bin/draft/add"
	// 获取草稿
	DraftGet = "https://api.weixin.qq.com/cgi-bin/draft/get"
	// 删除草稿
	DraftDelete = "https://api.weixin.qq.com/cgi-bin/draft/delete"
	// 更新草稿
	DraftUpdate = "https://api.weixin.qq.com/cgi-bin/draft/update"
	//count
	DraftCount = "https://api.weixin.qq.com/cgi-bin/draft/count"
	// 获取草稿列表
	DraftBatchGet = "https://api.weixin.qq.com/cgi-bin/draft/batchget"
)

const (
	EventPass    = "card_pass_check"     //通过审核
	EventNotPass = "card_not_pass_check" //未通过审核
)

// 数据分析
//https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html
const (
	//getusersummary
	GetUserSummary = "https://api.weixin.qq.com/datacube/getusersummary"
	//获取累计用户数据（getusercumulate）
	GetUserCumulate = "https://api.weixin.qq.com/datacube/getusercumulate"
	//获取图文群发每日数据（getarticlesummary）
	GetArticleSummary = "https://api.weixin.qq.com/datacube/getarticlesummary"
	//获取图文群发总数据（getarticletotal）
	GetArticleTotal = "https://api.weixin.qq.com/datacube/getarticletotal"
	//获取图文统计数据（getuserread）
	GetUserRead = "https://api.weixin.qq.com/datacube/getuserread"
	//获取图文统计分时数据（getuserreadhour）
	GetUserReadHour = "https://api.weixin.qq.com/datacube/getuserreadhour"
	//获取图文分享转发数据（getusershare）
	GetUserShare = "https://api.weixin.qq.com/datacube/getusershare"
	//获取图文分享转发分时数据（getusersharehour）
	GetUpstreamMsg          = "https://api.weixin.qq.com/datacube/getupstreammsg"
	GetUpstreamMsgHour      = "https://api.weixin.qq.com/datacube/getupstreammsghour"
	GetUpstreamMsgWeek      = "https://api.weixin.qq.com/datacube/getupstreammsgweek"
	GetUpstreamMsgMonth     = "https://api.weixin.qq.com/datacube/getupstreammsgmonth"
	GetUpstreamMsgDist      = "https://api.weixin.qq.com/datacube/getupstreammsgdist"
	GetUpstreamMsgDistWeek  = "https://api.weixin.qq.com/datacube/getupstreammsgdistweek"
	GetUpstreamMsgDistMonth = "https://api.weixin.qq.com/datacube/getupstreammsgdistmonth"
)

// 广告数据
const (
	// 获取公众号分广告位数据
	PublisherAdPosGeneral = "https://api.weixin.qq.com/publisher/stat?action=publisher_adpos_general"

	// 获取公众号返佣商品数据
	PublisherCpsGeneral = "https://api.weixin.qq.com/publisher/stat?action=publisher_cps_general"
	// 获取公众号结算收入数据及结算主体信息
	PublisherSettlement = "https://api.weixin.qq.com/publisher/stat?action=publisher_settlement"
)
