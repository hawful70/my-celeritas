package handlers

import (
	"net/http"

	celeritas "github.com/hawful70/my-celeritas"
)

type Handlers struct {
	App *celeritas.Celeritas
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, HomeHandler, nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) SessionTest(w http.ResponseWriter, r *http.Request) {
	myData := "foo"

	h.App.Session.Put(r.Context(), "bar", myData)

	err := h.App.Render.JetPage(w, r, SessionTestHandler, nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
