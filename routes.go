package luxe

import "regexp"

type Route struct {
	Method  string
	Path    string
	Handler HandlerFunc
	Regex   *regexp.Regexp
}

type Router struct{
	routes []Route
}

func NewRouter() *Router{
	return &Router{
		routes: make([]Route,0),
	}
}


