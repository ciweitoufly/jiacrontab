package jiaweb

import (
	"sync"
)

const (
	HTTPMethod_Any       = "ANY"
	HTTPMethod_GET       = "GET"
	HTTPMethod_POST      = "POST"
	HTTPMethod_PUT       = "PUT"
	HTTPMethod_DELETE    = "DELETE"
	HTTPMethod_PATCH     = "PATCH"
	HTTPMethod_HiJack    = "HIJACK"
	HTTPMethod_WebSocket = "WEBSOCKET"
	HTTPMethod_HEAD      = "HEAD"
	HTTPMethod_OPTIONS   = "OPTIONS"
)

type (
	HttpHandle func(httpCtx *HttpContext)
	Middleware func(httpCtx *HttpContext)
	Router     interface {
		ServeHttp(ctx *HttpContext)
		ServeFile(path, fileRoot string)
	}

	RouteNode interface {
		Use(m ...Middleware) *Node
		Middlewares() []Middleware
		Node() *Node
	}

	route struct {
		handleMap             map[string]HttpHandle
		NodeMap               map[string]*Node
		rwMutex               sync.RWMutex
		RedirectTrailingSlash bool
		RedirectFixedPath     bool
		HandleOPTIONS         bool
	}

	RouteHandle func(ctx *HttpContext)
)

var (
	SupportHTTPMethod map[string]bool
)

func NewRoute() *route {
	return &route{
		handleMap:             make(map[string]HttpHandle),
		NodeMap:               make(map[string]*Node),
		RedirectTrailingSlash: true,
		RedirectFixedPath:     true,
		HandleOPTIONS:         true,
	}
}

func (r *route) RegisterHandler(name string, handler HttpHandle) {
	r.rwMutex.Lock()
	r.handleMap[name] = handler
	r.rwMutex.Unlock()
}

func (r *route) GetHandler(name string) (HttpHandle, bool) {
	r.rwMutex.RLock()
	h, ok := r.handleMap[name]
	r.rwMutex.RUnlock()
	return h, ok
}

func (r *route) ServeHTTP(ctx *HttpContext) {
	req := ctx.Request().Request
	rw := ctx.Response().ResponseWriter
	path := req.URL.Path
	if root := r.NodeMap[req.Method]; root != nil {
		node, params, _ := root.getNode(path)
		if node.hander != nil {

		}
	}

}

func init() {
	SupportHTTPMethod[HTTPMethod_Any] = true
	SupportHTTPMethod[HTTPMethod_GET] = true
	SupportHTTPMethod[HTTPMethod_POST] = true
	SupportHTTPMethod[HTTPMethod_PUT] = true
	SupportHTTPMethod[HTTPMethod_DELETE] = true
	SupportHTTPMethod[HTTPMethod_PATCH] = true
	SupportHTTPMethod[HTTPMethod_HiJack] = true
	SupportHTTPMethod[HTTPMethod_WebSocket] = true
	SupportHTTPMethod[HTTPMethod_HEAD] = true
	SupportHTTPMethod[HTTPMethod_OPTIONS] = true
}
