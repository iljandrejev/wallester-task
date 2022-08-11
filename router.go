package main

import (
	"github.com/go-chi/chi"
	"sync"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	r := chi.NewRouter()

	customerController := ServiceContainer().InjectCustomerController()
	r.HandleFunc("/", customerController.List)
	r.MethodFunc("GET", "/create", customerController.Create)
	r.MethodFunc("GET", "/search", customerController.Search)
	r.MethodFunc("POST", "/create", customerController.Store)
	r.MethodFunc("GET", "/edit/{id}", customerController.Edit)
	r.MethodFunc("POST", "/edit/{id}", customerController.Update)
	r.MethodFunc("POST", "/delete/{id}", customerController.Delete)
	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
