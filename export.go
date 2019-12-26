package wxmini

import (
	"encoding/xml"
	"errors"

	"github.com/astaxie/beego"
	"github.com/visolleon/wxmini-sdk/common"
	"github.com/visolleon/wxmini-sdk/login"
	"github.com/visolleon/wxmini-sdk/pay"
	"github.com/visolleon/wxmini-sdk/subscribe"
)

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
		Secret: secret, // 微信应用APPId或小程序APPId
		MchID:  mchId,  // 商户号
		PayKey: payKey, // 支付密钥
	}
	return _instance
}

// 登录
func (this *AppProgram) Login(code, encryptedData, iv string) (wxUserInfo *login.WechatEncryptedData, err error) {
	wxConfig := &login.WxConfig{
		AppID:  beego.AppConfig.String("weChatAppID"),
		Secret: beego.AppConfig.String("weChatAppKey"),
	}
	return wxConfig.WexLogin(code, encryptedData, iv)
}

// 支付
func (this *AppProgram) Pay(openId string, money int, callbackURL, tradeNo, tradeInfo string) (results *pay.WaxPayRet, err error) {
	wePay := &pay.WePay{
		AppID:     this.AppID,
		MchID:     this.MchID,
		PayKey:    this.PayKey,
		NotifyURL: callbackURL, // 回调地址
		TradeType: "JSAPI",     // 小程序写"JSAPI",客户端写"APP"
		Body:      tradeInfo,   // 商品描述
	}
	results, err = wePay.WaxPay(money, openId, tradeNo) // 金额，以分为单位；open_id为获取的用户的open_id
	return results, err
}

// 解析微信支付回调数据
func (this *AppProgram) GetPayBackData(xmlData []byte) (*pay.WaxPayNotifyReq, error) {
	waxNotify := pay.WaxPayNotifyReq{}
	err := xml.Unmarshal(xmlData, &waxNotify)
	if err == nil {
		verifyParams := pay.WaxVerifyParams(waxNotify)
		valid := pay.WaxpayVerifySign(verifyParams, this.PayKey, waxNotify.Sign) //appKey 为自己在微信支付后台设置的API密钥

		if !valid {
			err = errors.New(common.ErrorPaySignInvalid)
		}
	}
	return &waxNotify, err
}

// 发布订阅消息
func (this *AppProgram) SendSubscribe(openId, templateId, jumpPage string, data map[string]interface{}) error {
	at, err := subscribe.GetAccessToken(this.AppID, this.Secret)
	if err == nil {
		err = subscribe.SendSubScribe(at.AccessToken, openId, templateId, jumpPage, data)
	}
	return err
}
