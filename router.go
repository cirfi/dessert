package dst

import (
	"net/http"
)

type route struct {
	handler  map[string]Handler
	pattern  string 			// TODO: "All", "Single", "RegExp"
}

type Router struct {
	method  map[string]*route
}

func NewRouter() *Router {
	r := new(Router)
	r.method = make(map[string]*route)
	r.method["GET"] = new(route)
	r.method["POST"] = new(route)
	r.method["OPTIONS"] = new(route)
	r.method["HEAD"] = new(route)
	r.method["PUT"] = new(route)
	r.method["DELETE"] = new(route)
	r.method["TRACE"] = new(route)
	return r
}

func (r *Router) RouterHandler(next Handler) Handler {
	fn := func(c *Context) {
		ok := r.match(c.Req.Method, c)
		if !ok {
			if c.Req.Method == "GET" {
				if next != nil {
					next(c)
				}
			} else {
				http.Error(c.Res, http.StatusText(404), 404)
			}
		}
	}
	return fn
}

func (r *Router) match(method string, c *Context) bool {
	path := pathFix(c.Req.URL.Path)
	fn, ok := r.method[method].handler[path]
	if ok == true {
		fn(c)
		return true
	} else {
		return false
	}
}

func (r *Router) addRouter(method string, path string, ptn string, handlers []RouterConstructor) {
	path = pathFix(path)
	if r.method[method].handler == nil {
		r.method[method].handler = make(map[string]Handler)
	}
	_, ok := r.method[method].handler[path]
	if ok != true {
		//
	}
	var handler Handler = nil
	for i := len(handlers) - 1; i >= 0; i-- {
		handler = handlers[i](handler)
	}
	r.method[method].handler[path] = handler
	r.method[method].pattern = ptn
}

func (r *Router) Get(path string, ptn string, handlers ...RouterConstructor) {
	r.addRouter("GET", path, ptn, handlers)
}

func (r *Router) Post(path string, ptn string, handlers ...RouterConstructor) {
	r.addRouter("POST", path, ptn, handlers)
}

func (r *Router) Options(path string, ptn string, handlers ...RouterConstructor) {
	r.addRouter("OPTIONS", path, ptn, handlers)
}

func (r *Router) Head(path string, ptn string, handlers ...RouterConstructor) {
	r.addRouter("HEAD", path, ptn, handlers)
}

func (r *Router) Put(path string, ptn string, handlers ...RouterConstructor) {
	r.addRouter("PUT", path, ptn, handlers)
}

func (r *Router) Delete(path string, ptn string, handlers ...RouterConstructor) {
	r.addRouter("DELETE", path, ptn, handlers)
}

func (r *Router) Trace(path string, ptn string, handlers ...RouterConstructor) {
	r.addRouter("TRACE", path, ptn, handlers)
}

func pathFix(path string) string {
	if path != "/" {
		if string(path[len(path) - 1]) == "/" {
			return path[:len(path) - 1]
		}
	}
	return path
}