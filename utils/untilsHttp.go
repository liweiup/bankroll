
package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

// Http Get请求基础函数, 通过封装Go语言Http请求, 支持火币网REST API的HTTP Get请求
// strUrl: 请求的URL
// strParams: string类型的请求参数, user=lxz&pwd=lxz
// return: 请求结果
func HttpGetRequest(strUrl string, mapParams,headerParams map[string]string) (string,error) {
	httpClient := &http.Client{}

	var strRequestUrl string
	if nil == mapParams {
		strRequestUrl = strUrl
	} else {
		strParams := Map2UrlQuery(mapParams)
		strRequestUrl = strUrl + "?" + strParams
	}
	// 构建Request, 并且按官方要求添加Http Header
	request, err := http.NewRequest("GET", strRequestUrl, nil)
	if nil != err {
		return "",err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	for k,v := range headerParams {
		request.Header.Add(k, v)
	}
	//request.Header.Add("Content-Type", "text/html; charset=utf8")
	// 发出请求
	response, err := httpClient.Do(request)
	if nil != err {
		return "",err
	}
	defer response.Body.Close()
	// 解析响应内容
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return "",err
	}
	resBody := string(body)
	contentType := response.Header.Get("Content-Type")
	if strings.Index(strings.ToLower(contentType),"gbk") != -1 {
		resBody = ConvertToString(&resBody,"gbk","utf-8")
	}
	if nil != err {
		return "",err
	}
	return resBody,nil
}
func HttpPostRequestBatchorder(strUrl string, mapParams interface{},headerParams map[string]string) (string,error) {
	httpClient := &http.Client{}
	jsonParams := ""
	switch reflect.TypeOf(mapParams).Kind() {
	case reflect.String:
		jsonParams = mapParams.(string)
	case reflect.Map:
		bytesParams, _ := json.Marshal(mapParams)
		jsonParams = string(bytesParams)
	}
	request, err := http.NewRequest("POST", strUrl, strings.NewReader(jsonParams))
	if nil != err {
		return "",err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept-Language", "zh-cn")
	for k,v := range headerParams {
		request.Header.Add(k, v)
	}
	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if nil != err {
		return "",err
	}
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return "",err
	}
	resBody := string(body)
	return resBody,nil
}


// 将map格式的请求参数转换为字符串格式的
// mapParams: map格式的参数键值对
// return: 查询字符串
func Map2UrlQuery(mapParams map[string]string) string {
	var strParams string
	for key, value := range mapParams {
		strParams += (key + "=" + value + "&")
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

// 将map格式的请求参数转换为字符串格式的,并按照Map的key升序排列
// mapParams: map格式的参数键值对
// return: 查询字符串
func Map2UrlQueryBySort(mapParams map[string]string) string {
	var keys []string
	for key := range mapParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var strParams string
	for _, key := range keys {
		strParams += key + "=" + mapParams[key] + "&"
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

// HMAC SHA256加密
// strMessage: 需要加密的信息
// strSecret: 密钥
// return: BASE64编码的密文
func ComputeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}