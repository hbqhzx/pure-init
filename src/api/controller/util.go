package controller

import (
	"errors"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	def "lib/db"
	"regexp"
	"time"
)

type State struct {
	actionAborted bool
	formParsed    bool
	query         *def.Query
	PayloadForm   *paramForm
}

type Context struct {
	*gin.Context
	*State
}

func (ctx *Context) textOutput(text string) {
	if !ctx.IsActionAborted() {
		ctx.String(200, text)
		ctx.AbortAction()
	}
}

func (ctx *Context) notFound() {
	if !ctx.IsActionAborted() {
		ctx.String(404, "Not found")
		ctx.AbortAction()
	}
}

func (ctx *Context) ShowPanicMsg(msg ...interface{}) {
	ctx.jsonFailure(msg...)
}

func (ctx *Context) jsonOutput(data interface{}) {
	origin := ctx.Request.Header.Get("Origin")
	allDomain := []string{
		"\\.jd\\.com(:\\d*)?$",
		"\\.jcloud\\.com(:\\d*)?$",
	}
	if !ctx.IsActionAborted() {
		for _, v := range allDomain {
			result, _ := regexp.MatchString(v, origin)
			if result {
				ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				headers := ctx.Request.Header.Get("Access-Control-Request-Headers")
				ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS, DELETE")
				ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
				if headers != "" {
					ctx.Writer.Header().Set("Access-Control-Allow-Headers", headers)
				}
			}
		}
	}
	ctx.JSON(200, data)
	ctx.AbortAction()
}

func (ctx *Context) traceInfo() gin.H {
	return gin.H{
		"id":        time.Now().UnixNano(),
		"timestamp": time.Now().Unix(),
		"srcIp":     ctx.ClientIP(),
		"destIp":    "",
	}
}

func (ctx *Context) jsonSuccess(ret ...interface{}) {
	if len(ret) > 0 {
		ctx.jsonOutput(gin.H{
			"trace": ctx.traceInfo(),
			"status": gin.H{
				"code": "OK",
				"msg":  "ok",
			},
			"data": gin.H{
				"item":  ret[0],
				"pager": ctx.GetPager(),
			},
		})
	} else {
		ctx.jsonOutput(gin.H{
			"trace": ctx.traceInfo(),
			"status": gin.H{
				"code": "OK",
				"msg":  "ok",
			},
		})
	}
}

func (ctx *Context) jsonError(err error) {
	ctx.jsonFailure(err.Error())
}

func (ctx *Context) jsonFailure(ret ...interface{}) {
	if len(ret) > 0 {
		ctx.jsonOutput(gin.H{
			"trace": ctx.traceInfo(),
			"status": gin.H{
				"code": "ClientError",
				"msg":  fmt.Sprintf("%s", ret[0]),
			},
		})
	} else {
		ctx.jsonOutput(gin.H{
			"trace": ctx.traceInfo(),
			"status": gin.H{
				"code": "ClientError",
				"msg":  "error",
			},
		})
	}
}

//正常情况不要是用这个
func (ctx *Context) DontUseJsonSuccess(ret ...interface{}) {
	ctx.jsonSuccess(ret...)
}

func CreateContext(gctx *gin.Context) *Context {
	ctx := &Context{gctx, &State{}}
	return ctx
}

func (state *State) IsActionAborted() bool {
	return state.actionAborted
}

func (state *State) AbortAction() {
	state.actionAborted = true
	panic(nil)
}

func (state *State) IsFormParsed() bool {
	return state.formParsed
}

func (state *State) SetFormParsed() {
	state.formParsed = true
}

func (state *State) SetQuery(query *def.Query) {
	state.query = query
}

func (state *State) GetQuery() *def.Query {
	return state.query
}

func (state *State) GetPager() *def.Pager {
	if state.query != nil {
		return state.query.Pager
	}
	return nil
}

func (ctx *Context) jsonInvalidParam() {
	ctx.jsonError(errors.New("Invalid param"))
}
