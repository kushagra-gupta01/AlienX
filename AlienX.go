package alienx

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
)

type ErrorHandler func(error, *Context) error

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	ctx      context.Context
}

func (c *Context) Render(component templ.Component) error {
	return component.Render(c.ctx, c.response)
}

type Handler func(c *Context) error

type AlienX struct {
	ErrorHandler ErrorHandler
	router       *httprouter.Router
}

func New() *AlienX {
	return &AlienX{
		router:       httprouter.New(),
		ErrorHandler: defaultErrorHandler,
	}
}

func (x *AlienX) Start(port string) error {
	return http.ListenAndServe(port, x.router)
}

func (x *AlienX) Get(path string, h Handler, plugs ...Handler) {
	x.router.GET(path, x.makeHTTPRouterHandle(h))
}

func (x *AlienX) makeHTTPRouterHandle(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := &Context{
			response: w,
			request:  r,
			ctx:      context.Background(),
		}
		if err := h(ctx); err != nil {
			//todo: handle the error from the error handler
			x.ErrorHandler(err, ctx)
		}
	}
}

func defaultErrorHandler(err error, c *Context) error {
	slog.Info("ERROR::", "err", err)
	return nil
}
