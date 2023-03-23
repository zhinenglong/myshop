package core

import (
	"github.com/gin-gonic/gin"
)

func GinServer(rootpath string) *gin.Engine {
	httpserver := gin.Default()
	httpserver.HTMLRender = LoadTemplates(rootpath + "/template")
	httpserver.MaxMultipartMemory = 80 * 1024 * 1024
	httpserver.Static("/static", rootpath)
	return httpserver
}

var GlobService *gin.Engine = nil

func init() {
	basepath := GetRootPath()
	GlobService = GinServer(basepath)
}
