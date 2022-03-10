package utils

import (
	"bankroll/config"
	"bankroll/global"
	"bankroll/global/redigo"
	"bankroll/service/model"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

const wxTokenKey = config.RedisKey + "WX:TOKEN"
type token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
}
func GetToken() string{
	url := global.Config.Else.WxPreurl + fmt.Sprintf("token?grant_type=client_credential&appid=%s&secret=%s",global.Config.Else.WxAppid,global.Config.Else.WxSecret)
	atoken, err := redigo.Dtype.String.Get(wxTokenKey).String()
	if atoken != "" {
		return atoken
	}
	request, err := HttpGetRequest(url,nil,nil)
	if err != nil {
		global.Zlog.Info("微信获取token错误：" + err.Error())
		return ""
	}
	var token token
	err = json.Unmarshal([]byte(request), &token)
	if err != nil {
		global.Zlog.Info("微信解析token错误：" + err.Error())
		return ""
	}
	redigo.Dtype.String.Set(wxTokenKey,token.AccessToken, token.ExpiresIn)
	return token.AccessToken
}

func SendMsg(token string,msgJsonArr []string) {
	atoken, err := redigo.Dtype.String.Get(wxTokenKey).String()
	if err != nil {
		global.Zlog.Info("微信消息推送错误：" + err.Error())
		return
	}
	url := global.Config.Else.WxPreurl + fmt.Sprintf("message/template/send?access_token=%s",atoken)
	for _, msgJson := range msgJsonArr {
		batchorder, err := HttpPostRequestBatchorder(url, msgJson, nil)
		if err != nil {
			continue
		}
		log.Println(batchorder)
	}
}

func GetModelMsg(Touser,TemplateID,URL,Value1,Value2,Value3 string) []string {
	toUserArr := strings.Split(Touser,",")
	wxPushDataArr := []string{}
	for _, v := range toUserArr {
		wxPushData := new(model.WxPushModel)
		wxPushData.Touser = v
		wxPushData.TemplateID = TemplateID
		wxPushData.URL = URL
		wxPushData.Data.First.Value = Value1
		wxPushData.Data.Keyword1.Value = Value2
		wxPushData.Data.Keyword2.Value = Value3
		marshal, _ := json.Marshal(wxPushData)
		wxPushDataArr = append(wxPushDataArr,string(marshal))
	}
	return wxPushDataArr
}