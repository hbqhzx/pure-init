package controller

func SayHi(ctx *Context) {
	ctx.jsonSuccess("hello,world!")

}
