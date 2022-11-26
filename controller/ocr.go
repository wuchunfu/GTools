package gtools

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gtools/configs"
	"gtools/internal"
	"gtools/util"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type bdOcrAccessToken struct {
	Token   string `json:"refresh_token"`
	Expires int32  `json:"expires_in"`
	Session string `json:"session_key"`
	Access  string `json:"access_token"`
	Scope   string `json:"scope"`
	Secret  string `json:"session_secret"`
}
type bdOcrResult struct {
	Words    []words `json:"words_result"`
	Num      string  `json:"words_result_num"`
	Id       string  `json:"log_id"`
	WordsStr string  `json:"words"`
}

type words struct {
	Words string `json:"words"`
}

func (a *App) GetBdOCRToken() (bool, string) {
	var host = configs.BdOcrTokenHost

	var param = map[string]string{
		"grant_type":    a.ConfigMap["bdocr"]["grantType"],
		"client_id":     a.ConfigMap["bdocr"]["clientId"],
		"client_secret": a.ConfigMap["bdocr"]["clientSecret"],
	}

	uri, err := url.Parse(host)
	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.BdOcrUrlParseErr, err.Error()))
		return false, err.Error()
	}
	query := uri.Query()
	for k, v := range param {
		query.Set(k, v)
	}
	uri.RawQuery = query.Encode()

	response, err := http.Get(uri.String())
	if err != nil {
		a.Log.Error(configs.BdOcrTokenGetErr, err.Error())
		return false, err.Error()
	}
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	if err != nil {
		a.Log.Error(configs.BdOcrTokenGetErr, err.Error())
		return false, err.Error()
	}
	resultStr := string(result)
	var newToken = bdOcrAccessToken{}
	err = json.Unmarshal([]byte(resultStr), &newToken)
	if err != nil {
		a.Log.Error(configs.BdOcrTokenGetErr, err.Error())
		return false, err.Error()
	}
	a.ConfigMap["bdocr"]["token"] = newToken.Access
	a.UpdateConfigItem(internal.ConfigItem {
		Name: "token",
		Type: "bdocr",
		Value: newToken.Access,
	})
	a.ConfigMap = a.GetConfigMap()
	return true, newToken.Access
}

func (a *App) BaiduOCR(ocrType int8) *util.Resp {
	// token := getOCRToken()
	var host string
	switch ocrType {
	case 0:
		host = configs.BdOcrGeneralBasic // 低精度 5W次/日
	case 1:
		host = configs.BdOcrAccurateBasic // 高精度 500次/日
	default:
		host = configs.BdOcrGeneralBasic
	}

	var accessToken = a.ConfigMap["bdocr"]["token"]
	if accessToken == "" {
		ok, token := a.GetBdOCRToken(); 
		if !ok {
			return util.Error(accessToken)
		} else {
			accessToken = token
		}
	}

	uri, err := url.Parse(host)
	if err != nil {
		a.Log.Error(configs.BdOcrUrlParseErr, err.Error())
		return util.Error(err.Error())
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()

	b, s := SavePngFromClipboard()
	if !b {
		return util.Error(s)
	}
	defer os.Remove(s)
	filebytes, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
	}

	image := base64.StdEncoding.EncodeToString(filebytes)
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", image)
	sendBody.Form.Add("language_type", "CHN_ENG")
	sendData := sendBody.Form.Encode()

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {

		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		a.ConfigMap["bdocr"]["token"] = ""
		a.Log.Error(configs.BdOcrRequestErr, err.Error())
		return util.Error(err.Error())
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		a.ConfigMap["bdocr"]["token"] = ""
		a.Log.Error(configs.BdOcrResponseErr, err.Error())
		return util.Error(err.Error())
	}
	resultJson := bdOcrResult{} 
	json.Unmarshal(result, &resultJson)
	wordsList := resultJson.Words
	var content = ""
	for _, v := range wordsList {
		content += v.Words + "\n"
	}
	return util.Success(content)
}
