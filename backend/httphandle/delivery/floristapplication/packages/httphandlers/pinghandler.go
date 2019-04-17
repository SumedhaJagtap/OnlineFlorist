package httphandlers

import (
	"log"
	"net/http"

	mthdroutr "OnlineFlorist/backend/httphandle/delivery/floristapplication/packages/mthdrouter"
	"OnlineFlorist/backend/httphandle/delivery/floristapplication/packages/resputl"
)

// PingHandler is a Basic ping utility for the service
type PingHandler struct {
	BaseHandler
}

func (p *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Got ping request: %v", r)
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

// Get function for PingHandler
func (p *PingHandler) Get(r *http.Request) resputl.SrvcRes {

	return resputl.Response200OK("OK")
}
