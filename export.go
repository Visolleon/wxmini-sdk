package wxmini

// 小程序应用实体
type AppProgram struct {
	AppID  string `json:"appid"`  // 微信APPID
	Secret string `json:"secret"` // 微信Secret
	MchID  string `json:"mchid"`  // 商户号
	PayKey string `json:"paykey"` // 支付密钥
}

// 创建新的对象
func NewAppProgram(appId, secret, mchId, payKey string) *AppProgram {
	_instance := &AppProgram{
		AppID:  appId,
		Secret: secret,
		MchID:  mchId,
		PayKey: payKey,
	}
	return _instance
}

// 登录
func (this *AppProgram) Login(code, encryptedData, iv string) {

}

// 支付
func (this *AppProgram) Pay(openId string, money int, callbackURL, tradeNo, tradeInfo string) {

}

// 发布订阅消息
func (this *AppProgram) Subscribe(openId, templateId, jumpPage string, data map[string]interface{}) {

}
