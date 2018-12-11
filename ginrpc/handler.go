package ginrpc

import "net/http"

// Handler call ServeHTTP for given next handler
type Handler struct {
	next http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.next.ServeHTTP(w, r)
}

//=======================================
