package configs

import "github.com/wailsapp/wails/v2/pkg/runtime"

var (
	LogFile = "%s/gtools.log"
	DBFile = "%s/gtools.db"
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
		Pattern: "*.md",
	}
)
