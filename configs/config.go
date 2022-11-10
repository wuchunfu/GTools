package configs

import "github.com/wailsapp/wails/v2/pkg/runtime"

var ConfigTypes = [...]string{"alioss", "localImgPath", "tuChuang"}
var AliOSS = [...]string{"point", "endPoint", "accessKeyId", "accessKeySecret", "bucketName", "projectDir"}
var LocalImgPath = [...]string{"path"}
var TuChuang = [...]string{"configType"}

var (
	LogFile = "%s/gtools.log"
	DBFile  = "%s/gtools.db"
)

// TODO 之后的阿里云数据默认为空值，并且数据入sqlite库
var (
	Point           = "oss-cn-beijing"
	Endpoint        = Point + ".aliyuncs.com"
	AccessKeyId     = "LTAI5tFNagDQNmkdeLoXrkdL"
	AccessKeySecret = "iUt3xEV5EgdmyLfASAwVNBsUfpQm4c"
	BucketName      = "ly-img-xiao"
	ProjectDir      = "test"
)

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
