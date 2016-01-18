package dst

import (
    //"fmt"
    "net/http"
)

func NotFoundPage(next Handler) Handler {
    return func(c *Context) {
        c.Res.Header().Set("Status", "404")
        http.Error(c.Res, http.StatusText(404), 404)
        //fmt.Fprintf(c.Res, "<h1>404 Not Found.</h1>, %s", c.Res)
        if next != nil {
            next(c)
        }
    }
}