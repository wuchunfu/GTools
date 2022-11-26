package configs

import "github.com/wailsapp/wails/v2/pkg/runtime"

// 存储所有配置项的类型[阿里云OSS，本地图片路径，图床类型，百度OCR，百度翻译，翻译类型]
var ConfigTypes = [...]string{"alioss", "limgpath", "imgbed", "bdocr", "bdtrans", "trans"}

// 按照类型存储所有配置项
var AliOSS = [...]string{"point", "endPoint", "accessKeyId", "accessKeySecret", "bucketName", "projectDir"}
var LocalImgPath = [...]string{"path"}
var ImgBed = [...]string{"configType"}
var BdOcr = [...]string{"grantType", "clientId", "clientSecret", "token"}
var BdTrans = [...]string{"appid", "secret", "salt", "from", "to"}
var Trans = [...]string{"transType"}

// 默认文件路径
var (
	LogFile = "%s/gtools.log"
	DBFile  = "%s/gtools.db"
)

// markdown 文件类型
var (
	MdFilter = runtime.FileFilter{
		DisplayName: "Markdown (*.md)",
		Pattern:     "*.md",
	}
	HtmlFilter = runtime.FileFilter{
		DisplayName: "HTML (*.html)",
		Pattern:     "*.html",
	}
)

// Mac webkit路径
var (
	WebkitPath = "%s/Library/Caches/com.wails.GTools/WebKit"
)

// 百度相关地址
var (
	BdOcrTokenHost     = "https://aip.baidubce.com/oauth/2.0/token"                // token
	BdOcrAccurateBasic = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic" // 通用文字识别-高精度版
	BdOcrGeneralBasic  = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"  // 通用文字识别-标准版

	BdTransUrl = "https://fanyi-api.baidu.com/api/trans/vip/translate"
)
