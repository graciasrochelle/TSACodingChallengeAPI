package main

import (
	"net/http"
)

type Chain struct {
	middleware middleware
	handlers   []chainHandler
}

type middleware struct {
	handler chainHandler
	next    *middleware
}

type chainHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type chainHandlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func NewChain() *Chain {
	return &Chain{NewMiddleware(), nil}
}

func (c *Chain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.middleware.ServeHTTP(w, r)
}

func (m middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.handler.ServeHTTP(w, r, m.next.ServeHTTP)
}

func (h chainHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h(w, r, next)
}

func (c *Chain) Use(h http.Handler) {
	if h != nil {
		ch := convert(h)
		c.UseChainHandler(ch)
	}
}

func (c *Chain) UseChainHandler(h chainHandler) {
	if h != nil {
		c.handlers = append(c.handlers, h)
	}
}

func (c *Chain) Build() {
	c.middleware = constructMiddleware(c.handlers)
}

func constructMiddleware(hs []chainHandler) middleware {
	var next middleware
	if len(hs) == 0 {
		return NewMiddleware()
	}
	next = constructMiddleware(hs[1:])
	return middleware{hs[0], &next}
}

func NewMiddleware() middleware {
	return middleware{
		chainHandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}),
		&middleware{},
	}
}

func convert(h http.Handler) chainHandler {
	return chainHandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		h.ServeHTTP(w, r)
		next(w, r)
	})
}
