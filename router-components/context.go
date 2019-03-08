package main

import "net/http"

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	router   *Router
	pnames   []string
	pvalues  []string
	index    int
	data     map[string]interface{}
	handlers []Handler
	writer   DataWriter
}
