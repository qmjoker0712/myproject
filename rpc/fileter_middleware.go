package rpc

import (
	"net/http"
	"time"
)

var timeFormat = "2006-01-02T15:04:05-0700"

// TODO 1. validate the service.method
// TODO 2.

func filterRequest(r *http.Request) bool {
	// All checks passed, create a codec that reads direct from the request body
	// untilEOF and writes the response to w and order the server to process a
	// single request.
	// ctx := r.Context()
	remoteAddr := r.RemoteAddr
	requestURI := r.RequestURI
	scheme := r.Proto
	local := r.Host
	userAgent := r.Header.Get("User-Agent")
	url := r.URL.RawQuery
	log.Debug("request info: remoteAddr=%v requestURI=%v scheme=%v host=%v time=%v", remoteAddr, requestURI, scheme, local, time.Now().Format(timeFormat))
	log.Debug("request info: userAgent=%#v url=%v", userAgent, url)
	return true
}
