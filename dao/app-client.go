package dao

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ReactAppRoot(basePath string, root *gin.RouterGroup) {
	root.GET("/app/*path", func(c *gin.Context) {
		var err error
		var info os.FileInfo
		var reqPath = c.Param("path")
		var srcPath = basePath + reqPath

		if info, err = os.Stat(srcPath); err != nil || info.IsDir() {
			srcPath = basePath + "/index.html"
		}
		http.ServeFile(c.Writer, c.Request, srcPath)
	})
}
