package main

type Router struct {
	RouteGroup
	IgnoreTrailingSlash bool
	UseEscapedPath bool

	pool sync.Pool
	routes []*Route
	namedRoutes map[string]*Route
	stores map[string]

	maxParams int
	notFound []Handler
	notFoundHandlers []Handler
}
