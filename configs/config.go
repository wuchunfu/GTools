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
	AccessKeyId     = "LTAI5tF8bKmV9Be9axkMgM5b"
	AccessKeySecret = "FrCuS7gqW1HW1zYVe3iKdnTLuEnmym"
	BucketName      = "ly-img-xiao"
	ProjectDir      = "test"
)

var (
	MdFilter = runtime.FileFilter{
		DisplayName: "Markdown (*.md)",
		Pattern: "*.md",
	}
)
