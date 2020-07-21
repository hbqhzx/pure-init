package controller

import (
	"fmt"
	def "lib/db"
	"lib/obj"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	PageSizeDefault = 30
	PageSizeMin     = 1
	PageSizeMax     = 30000
)

type R string // Required
type O string // Optional
type E string // Optional (padding with empty)

type paramForm struct {
	Form    url.Values
	Payload map[string]interface{}
}

func (form *paramForm) Get(key string) string {
	if val, ok := form.Payload[key]; ok {
		switch val.(type) {
		case int:
			str := strconv.Itoa(val.(int))
			return str
		case float64:
			return fmt.Sprintf("%.0f", val.(float64))
		}
		return val.(string)
	}
	return form.Form.Get(key)
}

func (ctx *Context) parseForm() *paramForm {
	if ctx.PayloadForm != nil {
		return ctx.PayloadForm
	}
	if !ctx.IsFormParsed() {
		ctx.Request.ParseForm()
		ctx.SetFormParsed()
	}

	form := paramForm{ctx.Request.Form, map[string]interface{}{}}
	payload := struct {
		Filter []struct {
			Name   string      `json:"name"`
			Method string      `json:"method"`
			Value  interface{} `json:"value"`
		} `json:"filter"`
	}{}

	for k, v := range ctx.Request.Form {
		if len(v) == 0 {
			form.Payload[k] = ""
		} else {
			form.Payload[k] = v[0]
		}
	}
	for _, v := range ctx.Request.Cookies() {
		form.Payload[v.Name] = v.Value
	}
	for _, v := range ctx.Params {
		form.Payload[v.Key] = v.Value
	}

	method := ctx.Request.Method
	if method == "POST" || method == "PUT" {
		var data interface{}
		ctx.Bind(&data)
		obj.Copy(&payload, &data)
		obj.Copy(&form.Payload, &data)
	}

	for _, v := range payload.Filter {
		switch v.Method {
		case ">", "gt":
			form.Payload[v.Name+":gt"] = v.Value
		case ">=", "ge":
			form.Payload[v.Name+":ge"] = v.Value
		case "<", "lt":
			form.Payload[v.Name+":lt"] = v.Value
		case "<=", "le":
			form.Payload[v.Name+":le"] = v.Value
		case "!=", "!", "ne", "not", "<>":
			form.Payload[v.Name+":ne"] = v.Value
		case "like", "%":
			form.Payload[v.Name+":like"] = v.Value
		case "prefix", "^":
			form.Payload[v.Name+":prefix"] = v.Value
		case "=", "", "eq", "in":
			fallthrough
		default:
			form.Payload[v.Name] = v.Value
		}
	}

	ctx.PayloadForm = &form
	return ctx.PayloadForm
}

/*************************************************************
 * Get request query 取请求参数
 *   ctx.query() 处理参数和分页、排序
 *   ctx.param() 只处理参数，不处理分页和排序
 *
 * Param: keys ...interface{}
 *   keys 取值范围如下：
 *     Required  必选参数：R("param1")  []R{"param2", "param3"}
 *     Optional  可选参数：O("param1")  []O{"param2", "param3"}  "param4"  []string{"param5", "param6"}
 *     Emptiable 自动补空：E("param1")  []E{"param2", "param3"}
 *
 * Return: *def.Query, error
 *
 * Example:
 *   query, err := ctx.query(R("r1"), []R{"r2", "r3"}, O("o1"), []O{"o2", "o3"}, []E{"e1", "e2"}, E("e3"), []string{"s1", "s2"}, "s3")
**************************************************************/
func (ctx *Context) query(keys ...interface{}) *def.Query {
	params := ctx.param(keys...)
	query := &def.Query{
		Param: &params,
	}
	ctx.queryCtrl(query)
	ctx.SetQuery(query)
	return query
}

func (ctx *Context) param(keys ...interface{}) map[string]interface{} {
	form := ctx.parseForm()
	params := make(map[string]interface{})
	for _, key := range keys {
		switch key.(type) {
		case R:
			v := form.Get(string(key.(R)))
			params[string(key.(R))] = v
			if v == "" {
				ctx.jsonFailure("Param `" + string(key.(R)) + "` is required")
			}
		case O:
			v := form.Get(string(key.(O)))
			if v != "" {
				params[string(key.(O))] = v
			}
		case string:
			v := form.Get(key.(string))
			if v != "" {
				params[key.(string)] = v
			}
		case E:
			params[string(key.(E))] = form.Get(string(key.(E)))
		case []R:
			for _, k := range key.([]R) {
				v := form.Get(string(k))
				params[string(k)] = v
				if v == "" {
					ctx.jsonFailure("Param `" + string(k) + "` is required")
				}
			}
		case []O:
			for _, k := range key.([]O) {
				v := form.Get(string(k))
				if v != "" {
					params[string(k)] = v
				}
			}
		case []string:
			for _, k := range key.([]string) {
				v := form.Get(k)
				if v != "" {
					params[k] = v
				}
			}
		case []E:
			for _, k := range key.([]E) {
				params[string(k)] = form.Get(string(k))
			}
		default:
			obj.Copy(key, &form.Payload)
			// ctx.Bind(key)
		}
	}
	return params
}

func (ctx *Context) queryCtrl(query *def.Query) {
	form := ctx.parseForm()

	if query.Pager == nil {
		query.Pager = &def.Pager{}
	}
	query.Pager.Page = 1
	query.Pager.PageSize = PageSizeDefault

	if page := toInt(form.Get("page")); page > 1 {
		query.Pager.Page = page
	}

	pageSize := toInt(form.Get("page_size"))
	if pageSize <= 0 {
		query.Pager.PageSize = PageSizeMax
	}

	if pageSize >= PageSizeMin && pageSize <= PageSizeMax {
		query.Pager.PageSize = pageSize
	}

	if matched, _ := regexp.MatchString("^\\w+$", form.Get("order_by")); matched {
		if query.Sorting == nil {
			query.Sorting = &def.Sorting{}
		}
		query.Sorting.OrderBy = form.Get("order_by")
		if strings.ToLower(form.Get("sort")) == "desc" {
			query.Sorting.Sort = "desc"
		} else {
			query.Sorting.Sort = "asc"
		}
	}
}

func toInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func (ctx *Context) paramInt(key string) int {
	return obj.ToInt(ctx.paramString(key))
}

func (ctx *Context) paramString(key string) string {
	form := ctx.parseForm().Payload
	val, ok := form[key]
	if !ok {
		return ""
	}
	switch val.(type) {
	case string:
		return val.(string)
	case int:
		return strconv.Itoa(val.(int))
	case float64:
		return fmt.Sprintf("%.0f", val.(float64))
	}
	return val.(string)
}
