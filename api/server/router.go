package server

import (
	"fmt"
	"net/http"
	"pure-init/api/controller"
	"pure-init/lib/log"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterFunc func(string, ...gin.HandlerFunc) gin.IRoutes
type HandleFunc func(*controller.Context)

type Router struct {
	Uri    string
	Method RouterFunc
	Action HandleFunc
}

var (
	router *gin.Engine
)

func initHttp() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Use(func(c *gin.Context) {
		start := time.Now()
		log.Infof("%#v", c.Request)
		c.Next()
		latency := time.Now().Sub(start)
		log.Infof("%15s %6s %3d %13v %s", c.ClientIP(), c.Request.Method, c.Writer.Status(), latency, c.Request.URL.Path)
	}, func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			ctx := controller.CreateContext(c)
			ctx.DontUseJsonSuccess()
		}
	})

	authReg := []string{
		"^/api/v1/app",
	}

	routerConfig := []Router{
		Router{"/hello", router.GET, controller.SayHi},
	}

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	auth := []*regexp.Regexp{}
	for _, reg := range authReg {
		auth = append(auth, regexp.MustCompile(reg))
	}

	needCheckAuth := func(uri string) bool {
		for _, reg := range auth {
			if reg.MatchString(uri) {
				return true
			}
		}
		return false
	}

	for _, config := range routerConfig {
		config.Method(config.Uri, makeHandler(config.Action, needCheckAuth(config.Uri)))
	}
}

func makeHandler(routerFunc HandleFunc, needCheckAuth bool) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx := controller.CreateContext(gctx)
		defer func() {
			if r := recover(); r != nil {
				ctx.ShowPanicMsg(r)
			}
		}()
		if needCheckAuth {
			// ctx.Auth()
		}
		routerFunc(ctx)
	}
}

func StartHttp(addr string) {
	initHttp()
	log.Info("init http succeed")
	router.Run(addr)
	fmt.Println(123)
}
