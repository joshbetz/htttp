package htttp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type Params map[string]string
type Handle func(*http.Request, Params) *Response
type Handler struct {
	httprouter *httprouter.Router
}

var access *log.Logger

func init() {
	access = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC)
}

func New() *Handler {
	return &Handler{
		httprouter: httprouter.New(),
	}
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Access logs
	access.Printf("[%s] [%s %s %s] [%s]", req.RemoteAddr, req.Method, req.URL, req.Host, req.UserAgent())

	// Default Access Control Headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")

	handler.httprouter.ServeHTTP(w, req)
}

func serve(w http.ResponseWriter, res *Response) {
	if res.Data != nil {
		switch (*res.Data).(type) {
		case string:
			if res.Code > 0 {
				w.WriteHeader(res.Code)
			}

			fmt.Fprint(w, (*res.Data).(string))
		default:
			w.Header().Set("Content-Type", "application/json")

			if res.Code > 0 {
				w.WriteHeader(res.Code)
			}

			if err := json.NewEncoder(w).Encode(res.Data); err != nil {
				log.Printf("Response error: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
	}
}
