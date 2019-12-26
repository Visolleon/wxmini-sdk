package subscribe

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/visolleon/wxmini-sdk/common"
	"github.com/visolleon/wxmini-sdk/utils"
)

type AccessTokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

// GetAccessToken
func GetAccessToken(appid, appsecret string) (*AccessTokenData, error) {
	if appid == "" {
		return nil, errors.New(common.ErrAppIDEmpty)
	}

	if appsecret == "" {
		return nil, errors.New(common.ErrAppIDEmpty)
	}

	params := url.Values{
		"grant_type": []string{"client_credential"},
		"appid":      []string{appid},
		"secret":     []string{appsecret},
	}

	body, err := utils.NewRequest("GET", common.CredentialAccessTokenURL, []byte(params.Encode()))
	if err != nil {
		return nil, err
	}

	m := new(AccessTokenData)
	err = json.Unmarshal(body, m)
	if err != nil {
		return m, err
	}
	if m.ErrMsg != "" {
		return m, errors.New(m.ErrMsg)
	}

	return m, nil
}

// SendSubScribe 发送订阅模版通知
// @accessToken: 令牌
// @userOpenId: 通知的用户OpenId
// @templateId: 通知模版Id
// @jumpPage: 跳转到小程序的页面路径
// @data: 模版数据
func SendSubScribe(accessToken, userOpenId, templateId, jumpPage string, data map[string]interface{}) error {
	params := map[string]interface{}{
		"access_token": accessToken,
		"touser":       userOpenId,
		"template_id":  templateId,
		"page":         jumpPage,
		"data":         data,
	}
	jsonData, err := json.Marshal(params)
	// log.Println(string(jsonData))
	body, err := utils.NewRequest("POST", fmt.Sprintf(common.SubScribeURL, accessToken), jsonData)
	if err != nil {
		return err
	}
	m := new(AccessTokenData)
	err = json.Unmarshal(body, m)
	if err != nil {
		return err
	}
	if m.ErrMsg != "" {
		return errors.New(m.ErrMsg)
	}

	return nil
}
