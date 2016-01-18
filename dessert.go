/* CirFi's Web Framework
 * Usage:
 * package main
 * import "github.com/cirfi/dessert"
 * func main() {
 * 	 server, _ := dst.DefaultServer()
 *   server.Run()
 * }
 */
package dst

import (
	"log"
	"net/http"
)

type Handler func(c *Context)

type Constructor func(http.Handler) http.Handler

type ContextConstructor func(Handler) http.Handler

type RouterConstructor func(Handler) Handler

type Dessert struct {
	port           string
	context        ContextConstructor
	router         RouterConstructor
	static         RouterConstructor
	staticPath     string
	notFound       RouterConstructor
	middlewares    []Constructor
	chainedHandler http.Handler
}

func (d *Dessert) Use(c Constructor) {
	d.middlewares = append(d.middlewares, c)
}

func (d *Dessert) UseContext(c ContextConstructor) {
	d.context = c
}

func (d *Dessert) UseRouter(r RouterConstructor) {
	d.router = r
}

func (d *Dessert) ServeStatic(r RouterConstructor, path string) {
	d.static = r
	d.staticPath = path
}

func (d *Dessert) NotFound(r RouterConstructor) {
	d.notFound = r
}

// Inspired by https://github.com/justinas/alice
func (d *Dessert) chain() {
	handler := d.context(d.router(d.static(d.notFound(nil))))
	for i := len(d.middlewares) - 1; i >= 0; i-- {
		handler = d.middlewares[i](handler)
	}
	d.chainedHandler = handler
}

func (d *Dessert) Run() {
	d.chain()
	//http.Handle("/", d.chainedHandler)
	log.Printf("The Dessert server is listening on port %v.", d.port[1:])
	http.ListenAndServe(d.port, d.chainedHandler)
}

func (d *Dessert) Port(p string) {
	d.port = ":" + p
}

func NewServer(port string) *Dessert {
	port = ":" + port
	return &Dessert{port: port}
}

func DefaultServer() (*Dessert, *Router) {
	d := NewServer("3000")
	d.Use(Logger)
	d.Use(Recover)
	d.UseContext(ContextHandler)
	r := NewRouter()
	d.UseRouter(r.RouterHandler)
	d.ServeStatic(d.Static, "")
	d.NotFound(NotFoundPage)
	return d, r
}
