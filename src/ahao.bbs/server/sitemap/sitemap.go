package sitemap

import (
	"ahao.bbs/server/common/uploader"
	"bytes"
	"compress/gzip"
	"strings"

	"github.com/mlogclub/simple/date"
	"time"

	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	"github.com/sirupsen/logrus"

	"ahao.bbs/server/common/urls"
	"ahao.bbs/server/config"
	"ahao.bbs/server/model"
	"ahao.bbs/server/model/constants"
)

const (
	changefreqAlways  = "always"
	changefreqHourly  = "hourly"
	changefreqDaily   = "daily"
	changefreqWeekly  = "weekly"
	changefreqMonthly = "monthly"
	changefreqYearly  = "yearly"
	changefreqNever   = "never"
)

var building = false

//Generate
func Generate() {
	if config.Instance.Env != "prod" {
		return
	}
	if building {
		logrus.Info("Sitemap in building...")
		return
	}
	building = true
	defer func() {
		building = false
	}()
	sm := stm.NewSitemap(0)
	sm.SetDefaultHost(config.Instance.BaseUrl)                  // 网站host
	sm.SetSitemapsHost(config.Instance.Uploader.AliyunOss.Host) // 上传到阿里云所以host设置为阿里云
	sm.SetSitemapsPath("sitemap2")                              // sitemap存放目录
	sm.SetVerbose(false)
	sm.SetPretty(false)
	sm.SetCompress(true)
	sm.SetAdapter(&myAdapter{})
	sm.Create()

	sm.Add(stm.URL{
		{"loc", urls.AbsUrl("/")},
		{"lastmod", time.Now()},
		{"changefreq", changefreqHourly},
	})

	sm.Add(stm.URL{
		{"loc", urls.AbsUrl("/topics")},
		{"lastmod", time.Now()},
		{"changefreq", changefreqHourly},
	})

	sm.Add(stm.URL{
		{"loc", urls.AbsUrl("/articles")},
		{"lastmod", time.Now()},
		{"changefreq", changefreqAlways},
	})

	sm.Add(stm.URL{
		{"loc", urls.AbsUrl("/projects")},
		{"lastmod", time.Now()},
		{"changefreq", changefreqDaily},
	})

}

type myAdapter struct {
}

func (adp *myAdapter) Bytes() [][]byte {
	return nil
}
func (adp *myAdapter) Write(loc *stm.Location, data []byte) {
	if stm.GzipPtn.MatchString(loc.Filename()) { // gzip
		var out []byte
		var in bytes.Buffer
		w := gzip.NewWriter(&in)
		_, _ = w.Write(data)
		_ = w.Close()
		out = in.Bytes()

		// 写入gzip格式数据
		adp.ossWrite(loc.PathInPublic(), out)

		// 写入原始数据
		adp.ossWrite(strings.ReplaceAll(loc.PathInPublic(), ".gz", ""), data)
	} else { // 非gzip
		adp.ossWrite(loc.PathInPublic(), data)
	}
}

//oss写入
func (adp *myAdapter) ossWrite(fileKey string, out []byte) {
	if url, err := uploader.PutObject(fileKey, out); err != nil {
		logrus.Error("Upload sitemap error:", err)
	} else {
		logrus.Info("Upload sitemap:", url)
	}

}
