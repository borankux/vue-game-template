package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)
import "github.com/gin-gonic/contrib/static"
func BindUI(engine *gin.Engine)  {
	engine.Use(static.Serve("/", static.LocalFile("ui/dist", false)))
	engine.NoRoute(func(ctx *gin.Context) {
		file,_ := ioutil.ReadFile("ui/dist/index.html")
		etag := fmt.Sprintf("%x", md5.Sum(file))

		ctx.Header("ETag", etag)
		ctx.Header("Cache-Control", "no-cache")

		if match := ctx.GetHeader("If-None-Match"); match != "" {
			if strings.Contains(match, etag) {
				ctx.Status(http.StatusNotModified)
				return
			}
		}

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", file)

	})
}
