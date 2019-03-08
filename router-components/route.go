package main

type Route struct {
	group          *RouteGroup
	method, path   string
	name, template string
	tags           []interface{}
	routes         []*Route
}
