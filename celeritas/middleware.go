package celeritas

import "net/http"

func (c *Celeritas) SessionLoad(next http.Handler) http.Handler {
	c.InfoLog.Println("Loading session for request")
	return c.Session.LoadAndSave(next)
}
