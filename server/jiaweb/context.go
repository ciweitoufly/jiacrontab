package jiaweb

type (
	Context interface {
		HttpServer() *HttpServer
		Response() *Response
		Request() *Request
		RouteNode() RouteNode
		Handler() HttpHandle
		RemoteIP() string
	}
	HttpContext struct {
		request    *Request
		response   *Response
		httpServer *HttpServer
		handler    HttpHandle
		routeNode  RouteNode
		params     map[string]string
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

func (ctx *HttpContext) RouteNode() RouteNode {
	return ctx.routeNode
}

func (ctx *HttpContext) Handler() HttpHandle {
	return ctx.handler
}

func (ctx *HttpContext) RemoteIP() string {
	return ctx.Request().RemoteIP()
}
