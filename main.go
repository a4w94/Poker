package main

import (
	apiserver_28gang "dbAPI/api/28gang"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	port      = ":5000"
	htmlroute = "template/html/*"
)

func main() {
	APIServerRun()

}

func APIServerRun() {
	router := InitRouter()
	fmt.Println("Route:", "localhost", port)
	s := &http.Server{
		Addr:           port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob(htmlroute)
	r.Static("/js", "./template/script")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		r.GET("/", StartHtml)
		r.GET("/28gang", apiserver_28gang.InitPage_28gang)
		//获取标签列表
		apiv1.GET("/ThmRate", apiserver_28gang.ThmRateAPI)
		//新建标签
		apiv1.POST("/CalThmRTP", apiserver_28gang.CalThmRTPAPI)

		apiv1.POST("/CalRangeRate", apiserver_28gang.CalRangeRateAPI)

	}
	return r
}

func StartHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "enter.html", nil)
}
