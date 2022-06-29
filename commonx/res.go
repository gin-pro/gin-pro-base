package commonx

import "fmt"

const (
	ResponseCodeOk          = 0 //请求成功
	ResponseCodeAuth        = 2 //无权限
	ResponseCodeErr         = 3 //请求失败
	ResponseCodeErrs        = 4 //请求失败 args
	ResponseCodeErrParam    = 5 //参数错误
	ResponseCodeErrNotFound = 6 //未找到数据

	// install
	ResponseCodeInstallWait = 100 //等待数据库

	// sys 10000+
	ResponseCodeSysNoSetting = 10001
	// login 20000+
	ResponseCodeLoginNoSetOAauth = 20001
)

var msgs map[string]string

func GetResponseMsg(lc string, code int, args ...interface{}) string {
	if msgs == nil {
		msgs = make(map[string]string)
		initZHUiMsg()
	}
	fs := msgs[fmt.Sprintf("%s:%d", lc, code)]
	if len(args) > 0 {
		return fmt.Sprintf(fs, args...)
	}
	return fs
}

func setResponseMsg(lc string, code int, msgfmt string) {
	msgs[fmt.Sprintf("%s:%d", lc, code)] = msgfmt
}
func initZHUiMsg() {
	lc := "zh"
	setResponseMsg(lc, ResponseCodeOk, "请求成功")
	setResponseMsg(lc, ResponseCodeAuth, "无权限")
	setResponseMsg(lc, ResponseCodeErr, "请求失败")
	setResponseMsg(lc, ResponseCodeErrs, "请求失败:%v")
	setResponseMsg(lc, ResponseCodeErrParam, "参数错误")
	setResponseMsg(lc, ResponseCodeErrNotFound, "尚未找到任何内容")

	setResponseMsg(lc, ResponseCodeLoginNoSetOAauth, "未配置oauth信息")
}
