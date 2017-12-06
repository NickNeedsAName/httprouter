package httprouter

import (
	"context"
	"net/http"
)

type paramsKey struct{}

var contextParamsKey = paramsKey{}

func (r *Router) Handler(method, path string, handler http.Handler) {
	r.Handle(method, path,
		func(w http.ResponseWriter, req *http.Request, p Params) {
			ctx := req.Context()
			ctx = context.WithValue(ctx, contextParamsKey, p)
			req = req.WithContext(ctx)
			handler.ServeHTTP(w, req)
		},
	)
}

func ParamsFromContext(ctx context.Context) Params {
	p, _ := ctx.Value(contextParamsKey).(Params)
	return p
}
