package controller

import (
	"net/http"
)

// RedirectTLS to port 443
func RedirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://abinash.tech:443/"+r.RequestURI, http.StatusMovedPermanently)
}
