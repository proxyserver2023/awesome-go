package main

type RouteGroup struct {
	prefix   string
	router   *Router
	handlers []Handler
}
