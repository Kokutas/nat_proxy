package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Web(webAddress string) {
	router := gin.Default()

	// 静态资源加载，本例为css,js以及资源图片
	router.Static("/static", "./web/static/*")
	router.StaticFile("/favicon.ico", "./web/static/favicon.ico")
	router.LoadHTMLGlob("./web/static/templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "内网穿透首页",
			"ip":"123",
		})
	})
	// 获取本机ip
	// TODO: 比较好 http://api.k780.com/?app=ip.local&format=json
	// {"success":"1","result":{"ip":"115.196.249.101","proxy":"1","att":"中国,浙江,杭州","operators":"电信"}} EOF
	// http://ip.360.cn/IPShare/info
	// http://www.taobao.com/help/getip.php
	// https://cloud.tencent.com/developer/article/1152362
	router.Run(webAddress)
}
