package main

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"

	"ahao.bbs/server/services"
	"ahao.bbs/server/sitemap"
)

func startSchedule() {
	c := cron.New()
	addCronFunc(c, "@every 30m", func() {
		services.ArticleService.GenerateRss()
		services.TopicSerci
	})

	addCornFunc(c, "0 0 4 ? * *", func() {
		sitemap.Generate()
	})
	c.Start()
}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}
