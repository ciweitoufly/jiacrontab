package main

import (
	"net/http"
	"jiacrontab/server/view"
)


func filterReq(rw http.ResponseWriter, r *http.Request, m *view.ModelView) bool {
	if r.URL.Path == "/favicon.ico" {
		return false
	}
	m.Locals("action", r.URL.Path)


	if err := globalReqFilter.filter(rw, r); err != nil {
		m.RenderHtml2([]string{"public/error"}, map[string]interface{}{
			"error": err.Error(),
		}, nil)
		return false
	} else {
		return true
	}
}

func omit(rw http.ResponseWriter, r *http.Request, m *view.ModelView) bool {
	if r.URL.Path == "/favicon.ico" {
		return false
	} else {
		return true
	}
}

func checkLogin(rw http.ResponseWriter, r *http.Request, m *view.ModelView) bool {

	var userinfo map[string]interface{}
	ok := globalJwt.auth(rw, r, &userinfo)
	if ok {
		for k, v := range userinfo {
			m.ShareData(k, v)
			m.Locals(k, v)
		}
	}
	if !ok && r.URL.Path == "/login" {
		ok = true
	}

	if !ok {
		http.Redirect(rw, r, "/login", http.StatusFound)
	}
	return ok
}
