package main

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

//NewMiddleWareHandler ...
func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		defer func() { recover() }()
		log.Panicln("Too many requests")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

//RegisterHandlers ...
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/ce", ce)
	return router
}

func ce(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//为了演示效果这块设置了等待
	time.Sleep(time.Second * 100)
	t, _ := template.ParseFiles("./videos/ce.html")
	t.Execute(w, nil)
}

//流量控
func main() {
	r := RegisterHandlers()
	//里面的参数2为设置的最大流量
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}
