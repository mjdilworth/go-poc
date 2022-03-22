package handlers

import (
	"net/http"
)

//placeholder
func verifyUserPass(user string, pass string) bool {
	return true
}

func Health(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    w.Write([]byte(`{"message": "I am healthy"}`))
}

func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
    w.Header().Add("Content-Type", "application/json")
    w.Write([]byte(`{"message": "HTTPS Served by GO"}`))
}

func Auth(w http.ResponseWriter, req *http.Request) {
	user, pass, ok := req.BasicAuth()
	if ok && verifyUserPass(user, pass) {
		w.Write([]byte(`{"message": "You get to see the secret"}`))
		//fmt.Fprintf(w, "You get to see the secret\n")
	} else {
		// i should redirect to login page
		w.Header().Set("WWW-Authenticate", `Basic realm="api"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}