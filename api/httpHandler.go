package api

import "net/http"

type httpHandler struct {
	handlers map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewHttpHandler() httpHandler {
	return httpHandler{
		handlers: make(map[string]func(w http.ResponseWriter, r *http.Request)),
	}
}

func (httpHandler httpHandler) AddEndpoint(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if otherHandlers, ok := httpHandler.handlers[pattern]; ok {
		httpHandler.handlers[pattern] = func(w http.ResponseWriter, r *http.Request) {
			if r.Method == method {
				handler(w, r)
			} else {
				otherHandlers(w, r)
			}
		}
	} else {
		httpHandler.handlers[pattern] = func(w http.ResponseWriter, r *http.Request) {
			if r.Method == method {
				handler(w, r)
			}
		}
	}
}

func (httpHandler httpHandler) ListenAndServe(port string) {
	for pattern, handler := range httpHandler.handlers {
		http.HandleFunc(pattern, handler)
	}
	http.ListenAndServe(port, nil)
}
