package jiaweb

type (
	Context interface {
	}
	HttpContext struct {
		request    *Request
		response   *Response
		httpServer *HttpServer
	}
)

func (ctx *HttpContext) reset(r *Request, rw *Response, httpServer *HttpServer) {
	ctx.request = r
	ctx.response = rw
	ctx.httpServer = httpServer
}

func (ctx *HttpContext) Request() *Request {
	return ctx.request
}

func (ctx *HttpContext) Response() *Response {
	return ctx.response
}
