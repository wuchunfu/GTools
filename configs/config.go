package configs

import "github.com/wailsapp/wails/v2/pkg/runtime"

// 存储所有配置项的类型
var ConfigTypes = [...]string{"alioss", "localImgPath", "imgBed"}
// 按照类型存储所有配置项
var AliOSS = [...]string{"point", "endPoint", "accessKeyId", "accessKeySecret", "bucketName", "projectDir"}
var LocalImgPath = [...]string{"path"}
var ImgBed = [...]string{"configType"}

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
)

// Mac webkit路径
var (
	WebkitPath = "%s/Library/Caches/com.wails.GTools/WebKit"
)
