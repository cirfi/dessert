package dst

import (
    "net/http"
    "os"
)

func Static(next Handler, staticPath string) Handler {
    return func(c *Context) {
        const indexPage = "/index.html"
        path := staticPath + c.Req.URL.Path[1:]
        f, err := os.Stat(path)
        if err == nil {
            if f.IsDir() {
                _, err := os.Stat(pathFix(path) + indexPage)
                if err == nil {
                    http.ServeFile(c.Res, c.Req, path)
                } else {
                    if next != nil {
                        next(c)
                    }
                }
            } else {
                http.ServeFile(c.Res, c.Req, path)
            }
        } else {
            if next != nil {
                next(c)
            }
        }
        //http.FileServer((http.Dir(""))).ServeHTTP(c.Res, c.Req)
    }
}
