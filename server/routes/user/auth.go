package user

import (
	"crypto/md5"
	"fmt"
	"jiacrontab/server/view"
	"net/http"
)

func Login(rw http.ResponseWriter, r *http.Request, m *view.ModelView) {
	if r.Method == http.MethodPost {

		u := r.FormValue("username")
		pwd := r.FormValue("passwd")
		remb := r.FormValue("remember")

		if u == globalConfig.User && pwd == globalConfig.Passwd {
			md5p := fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
			if remb == "yes" {
				globalJwt.accessToken(rw, r, u, md5p)
			} else {
				globalJwt.accessTempToken(rw, r, u, md5p)
			}

			http.Redirect(rw, r, "/", http.StatusFound)
			return
		}

		m.RenderHtml2([]string{"public/error"}, map[string]interface{}{
			"error": "auth failed",
		}, nil)

	} else {
		var user map[string]interface{}
		if globalJwt.auth(rw, r, &user) {
			http.Redirect(rw, r, "/", http.StatusFound)
			return
		}
		m.RenderHtml2([]string{"login"}, nil, nil)

	}
}
