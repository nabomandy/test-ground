package router

import (
	"net/http"

	dao "dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/dao"
	v1 "dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/process/api/v1"
	apiTest "dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/process/api/v1/apiTest"
	dbData "dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/process/api/v1/dbData"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// 메인 라우터
func InitRouter() *gin.Engine {
	r := gin.Default()

	// /vurix-dms/api
	// /vurix-dms/app

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})
	// /base/app 설정을 위해서
	dao.ReactAppRoot("./client/build", r.Group("vurix-dms"))

	subRg := r.Group("vurix-dms/api")
	{
		subRg.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		subRg.Static("/docs", "./docs")

		subRg.GET("/version", Version)
	}
	InitV1Router(subRg)

	// 서비스 모니터링용 pprof 초기화
	pprof.RouteRegister(subRg, "pprof")

	return r
}

// 메인 라우터의 버전
func Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": "1.0.0",
	})
}

// V1 라우팅 처리
func InitV1Router(rg *gin.RouterGroup) {
	rg.Use()
	{
		rgV1 := rg.Group("/v1")
		{
			rgV1.GET("/version", v1.Version)
		}
		InitTestRouter(rgV1)
		InitDBRouter(rgV1)
	}
}

// V1 하위 DB 관련 라우팅
func InitDBRouter(rg *gin.RouterGroup) {
	rg.Use()
	{
		rgSub := rg.Group("/dbData")
		{
			rgSub.POST("/getDataList", dbData.GetDataList)
			// rgSub.POST("/getDataList", dbData.GetDataList2)
		}
	}
}

// V1 하위 샘플 API 라우팅
func InitTestRouter(rg *gin.RouterGroup) {
	rg.Use()
	{
		rgSub := rg.Group("/apiTest")
		{
			rgSub.GET("/getAPIData", apiTest.GetAPIData)
			rgSub.GET("/getAPIData2", apiTest.GetAPIData2)
			rgSub.POST("/postAPIData", apiTest.PostAPIData)
		}
	}
}
