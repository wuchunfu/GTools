package gtools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"gtools/configs"
	"gtools/util"
	"io"
	"net/http"
	"strings"
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
