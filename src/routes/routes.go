package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	appConf "github.com/kukkar/bussinesscard/conf"
	controller "github.com/kukkar/bussinesscard/src/controller"
	"github.com/kukkar/common-golang/pkg/middleware"
)

func Routes(route *gin.Engine) {

	gConf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}
	appConfig, err := appConf.GetAppConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(appConfig)
	v1 := route.Group(string(gConf.AppName) + "/v1")
	{
		defaultMiddleware := middleware.DefaultMiddleware{}
		v1.GET("/hellworld", defaultMiddleware.MonitorRequest(), controller.HelloWorld)
	}
}
