package server

import (
	"net/http"
	"net/http/httputil"
	"pure-init/api/controller"
	"regexp"
	"runtime"
	"time"

	"github.com/cihub/seelog"
	"gopkg.in/gin-gonic/gin.v1"
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
		seelog.Tracef("%#v", c.Request)
		c.Next()
		latency := time.Now().Sub(start)
		seelog.Infof("%15s %6s %3d %13v %s", c.ClientIP(), c.Request.Method, c.Writer.Status(), latency, c.Request.URL.Path)
	}, func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			ctx := controller.CreateContext(c)
			ctx.DontUseJsonSuccess()
		}
	})

	authReg := []string{
		"^/api/v2/app",
		"^/api/v2/group",
		"^/api/v2/token",
		"^/api/v3/app",
	}

	routerConfig := []Router{

		//Router{"/hello", router.GET, controller.Testtsdb},
		Router{"/hello", router.GET, controller.SayHi},
	}

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
				stack := make([]byte, 1024*8)
				stack = stack[:runtime.Stack(stack, false)]
				httprequest, _ := httputil.DumpRequest(ctx.Request, true)
				seelog.Trace(string(stack))
				seelog.Trace(string(httprequest))
				//panic(r)
				ctx.ShowPanicMsg(r)
			}
		}()
		if needCheckAuth {
			//ctx.Auth()
		}
		routerFunc(ctx)
	}
}

func StartHttp(addr string) {
	initHttp()
	router.Run(addr)
}