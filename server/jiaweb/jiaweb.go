package jiaweb

import (
	"jiacrontab/server/jiaweb/logger"
)

type (
	JiaWeb struct {
		HttpServer              *HttpServer
		Middlewares             []Middleware
		ExceptionHandler        ExceptionHandle
		NotFoundHandler         StandardHandle
		MethodNotAllowedHandler StandardHandle
	}

	// 自定义异常处理
	ExceptionHandle func(Context, error)

	StandardHandle func(Context)
	HttpHandle     func(httpCtx *HttpContext)
)

const (
	DefaultHTTPPort    = 8080
	RunModeDevelopment = "development"
	RunModeProduction  = "production"
)

func New() *JiaWeb {
	app := &JiaWeb{
		HttpServer:  NewHttpServer(),
		Middlewares: make([]Middleware, 0),
	}

	logger.InitJiaLog()

	return app
}

func Classic() *JiaWeb {
	app := New()

}

func (app *JiaWeb) SetEnableLog(enableLog bool) {
	logger.SetEnableLog(enableLog)

}

func (app *JiaWeb) SetLogPath(path string) {
	logger.SetLogPath(path)
}
