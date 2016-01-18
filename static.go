package dst

import (
    "net/http"
    "os"
)

func (d *Dessert) Static(next Handler) Handler {
    return func(c *Context) {
        const indexPage = "/index.html"
        path := c.Req.URL.Path[1:]
        f, err := os.Stat(d.staticPath + path)
        if err == nil {
            if f.IsDir() {
                _, err := os.Stat(pathFix(path) + indexPage)
                if err == nil {
                    http.ServeFile(c.Res, c.Req, d.staticPath + path)
                } else {
                    if next != nil {
                        next(c)
                    }
                }
            } else {
                http.ServeFile(c.Res, c.Req, d.staticPath + path)
            }
        } else {
            if next != nil {
                next(c)
            }
        }
        //http.FileServer((http.Dir(""))).ServeHTTP(c.Res, c.Req)
    }
}
