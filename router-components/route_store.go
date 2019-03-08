package main

type routeStore interface {
	Add(key string, data interface{}) int
	Get(key string, pvalues []string) (data interface{}, pnames []string)
	String() string
}
