package luxe

type Route struct {
	Method  string
	Path    string
	Handler HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{
		routes: make([]Route, 0),
	}
}

func (r *Router) addRoute(method, path string, handler HandlerFunc) {
	route := Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
	r.routes = append(r.routes, route)

}
func (r *Router) GET(path string, handler HandlerFunc) {
	r.addRoute(METHODGET.ToString(), path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.addRoute(METHODPOST.ToString(), path, handler)
}

func (r *Router) PUT(path string, handler HandlerFunc) {
	r.addRoute(METHODPUT.ToString(), path, handler)
}

func (r *Router) DELETE(path string, handler HandlerFunc) {
	r.addRoute(METHODDELETE.ToString(), path, handler)
}

func (r *Router) PATCH(path string, handler HandlerFunc) {
	r.addRoute(METHODPATCH.ToString(), path, handler)
}

func (r *Router) HandleRequest(l *LTX) {
	method := l.GetMethod()
	path := l.GetPath()

	for _, route := range r.routes {
		if route.Method != method {
			continue
		}

		// Check exact path match first
		if route.Path == path {
			route.Handler(l)
			return
		}
	}
	l.SendJSON(404, H{
		"error":  "Route not found",
		"path":   path,
		"method": method,
	})
}

// GET adds a GET route
func (l *Luxe) GET(path string, handler HandlerFunc) {
	l.router.GET(path, handler)
}

// POST adds a POST route
func (l *Luxe) POST(path string, handler HandlerFunc) {
	l.router.POST(path, handler)
}

// PUT adds a PUT route
func (l *Luxe) PUT(path string, handler HandlerFunc) {
	l.router.PUT(path, handler)
}

// DELETE adds a DELETE route
func (l *Luxe) DELETE(path string, handler HandlerFunc) {
	l.router.DELETE(path, handler)
}

// PATCH adds a PATCH route
func (l *Luxe) PATCH(path string, handler HandlerFunc) {
	l.router.PATCH(path, handler)
}
