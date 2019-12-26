package common

const (
	//https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
	// AccessTokenURL code获取access_token
	AccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token"

	// RefreshTokenURL 重新获取access_token
	RefreshTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token"

	// UserInfoURL 通过access_token获取userInfo
	UserInfoURL = "https://api.weixin.qq.com/sns/userinfo"

	// UnifiedOrderURL 微信统一下单
	// 文档地址:
	// https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1419317853&token=&lang=zh_CN
	UnifiedOrderURL = "https://api.mch.weixin.qq.com/pay/unifiedorder"

	// CheckAccessTokenURL 检验授权凭证（access_token）是否有效
	CheckAccessTokenURL = "https://api.weixin.qq.com/sns/auth"

	// JsCode2SessionURL 临时登录凭证校验接口
	JsCode2SessionURL = "https://api.weixin.qq.com/sns/jscode2session"

	// SendRedPackURL 发送现金红包
	SendRedPackURL = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"

	// "?grant_type=client_credential&appid=APPID&secret=APPSECRET"
	// CredentialAccessTokenURL 获取access_token令牌
	CredentialAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token"

	// ?access_token=ACCESS_TOKEN
	// 模SubScribeURL 板订阅消息
	SubScribeURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"
)
