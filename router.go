package htttp

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (handler *Handler) Get(path string, handle Handle) {
	handler.Handle("HEAD", path, handle)
	handler.Handle("GET", path, handle)
}

func (handler *Handler) Post(path string, handle Handle) {
	handler.Handle("POST", path, handle)
}

func (handler *Handler) Put(path string, handle Handle) {
	handler.Handle("PUT", path, handle)
}

func (handler *Handler) Patch(path string, handle Handle) {
	handler.Handle("PATCH", path, handle)
}

func (handler *Handler) Delete(path string, handle Handle) {
	handler.Handle("DELETE", path, handle)
}

func (handler *Handler) Handle(method, path string, handle Handle) {
	handler.Router.Handle(method, path, func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		params := map[string]string{}
		for _, p := range ps {
			params[p.Key] = p.Value
		}

		res := handle(req, params)
		serve(w, res)
	})
}
