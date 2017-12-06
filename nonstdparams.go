package httprouter

import "net/http"

func (r *Router) Handler(method, path string, handler http.Handler) {
	r.Handle(method, path,
		func(w http.ResponseWriter, req *http.Request, _ Params) {
			handler.ServeHTTP(w, req)
		})
}
