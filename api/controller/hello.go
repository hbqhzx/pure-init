package controller

// @Summary 服务健康状态
// @Tags hello
// @Produce application/json
// @Success 200 object dao.GetHelloRes
// @Failure 400 object dao.Response
// @Router /hello [get]
func SayHi(ctx *Context) {
	ctx.jsonSuccess("hello,world!")
}
