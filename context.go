package dst

import (
    "net/http"
)

type Context struct {
    store map[interface{}]interface{}
    Res http.ResponseWriter
    Req *http.Request
}

func (c *Context) Get(key interface{}) interface{} {
    value, ok := c.store[key]
    if ok == true {
        return value
    }
    return nil
}

func (c *Context) Set(key, value interface{}) {
    c.store[key] = value
}

func ContextHandler(next Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        c := new(Context)
        c.store = make(map[interface{}]interface{})
        c.Res = w
        c.Req = r
        if next != nil {
            next(c)
        }
    }
    return http.HandlerFunc(fn)
}