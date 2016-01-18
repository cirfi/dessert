package main

import (
	"github.com/cirfi/dessert"
	"fmt"
)

func main() {
	d, r := dst.DefaultServer()
	r.Get("/", "All", func(next dst.Handler) dst.Handler {
		return func(c *dst.Context) {
			c.Set("Test", "Nothing")
			if next != nil {
				next(c)
			}
		}
	}, func(next dst.Handler) dst.Handler {
		return func(c *dst.Context) {
			fmt.Fprintf(c.Res, "%v\n", c.Get("Test"))
			if next != nil {
				next(c)
			}
		}
	})
	d.Run()
}
