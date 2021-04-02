package common

import (
	"ahao.bbs/server/common/html"
	"ahao.bbs/server/common/markdown"
	"ahao.bbs/server/config"
	"ahao.bbs/server/model/constants"
	"github.com/mlogclub/simple"
	"math/rand"
	"strconv"
)

//是否是正式环境
func IsProd() bool {
	return config.Instance.Env == "prod"
}
func GetSummary(contentType string, content string) (summary string) {
	if contentType == constants.ContentTypeMarkdown {
		summary = markdown.GetSummary(content, constants.SummaryLen)
	} else if contentType == constants.ContentTypeHtml {
		summary = html.GetSummary(content, constants.SummaryLen)
	} else {
		summary = simple.GetSummary(content, constants.SummaryLen)
	}
	return
}

//生成随机验证码
func RandomCode(len int) string {
	if len < 4 {
		len = 4
	}
	var code string
	for i := 0; i < len; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}
