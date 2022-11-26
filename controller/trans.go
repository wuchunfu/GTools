package gtools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gtools/configs"
	"gtools/util"
	"io"
	"net/http"
	"strings"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

type BdTransResult struct {
	From        string      `json:"from"`
	To          string      `json:"to"`
	TransResult []transItem `json:"trans_result"`
}
type transItem struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

func (a *App) BaiDuTrans(query string) *util.Resp {
	var appid = a.ConfigMap["bdtrans"]["appid"]
	var salt = a.ConfigMap["bdtrans"]["salt"]
	var secret = a.ConfigMap["bdtrans"]["secret"]
	var str = appid + query + salt + secret
	h := md5.New()
	h.Write([]byte(str))
	var sign = hex.EncodeToString(h.Sum(nil))
	req, _ := http.NewRequest("GET", configs.BdTransUrl, nil)
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("from", "auto")
	q.Add("to", "zh")
	q.Add("appid", appid)
	q.Add("salt", salt)
	q.Add("sign", sign)
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		a.Log.Error(configs.BdTransRequestErr, err.Error())
		return util.Error("翻译失败，请检查百度翻译设置项")
	}
	defer resp.Body.Close()
	var result []byte
	result, _ = io.ReadAll(resp.Body)
	if strings.Contains(string(result), "error_code") {
		return util.Error("翻译失败，请检查百度翻译设置项")
	}
	var resultJson = new(BdTransResult)
	err = json.Unmarshal(result, &resultJson)
	if err != nil {
		return util.Error("翻译结果解析失败，请重试")
	}
	var resultStr = ""
	for _, v := range resultJson.TransResult {
		resultStr += (v.Dst + "\n")
	}
	return util.Success(resultStr)
}

func (a *App)TxTrans() {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		"AKIDgASWR9fy3tVu4VigSzHLEsJVTE0Xq0xh",
		"aw61iBPdfU5XmodFguuhIhuukPrDMK6e",
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tmt.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := tmt.NewClient(credential, "ap-beijing", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := tmt.NewTextTranslateRequest()

	request.SourceText = common.StringPtr("test")
	request.Source = common.StringPtr("auto")
	request.Target = common.StringPtr("zh")
	request.ProjectId = common.Int64Ptr(0)

	// 返回的resp是一个TextTranslateResponse的实例，与请求对象对应
	response, err := client.TextTranslate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
