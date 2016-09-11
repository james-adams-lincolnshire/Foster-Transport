package handlers

import (
	"github.com/golang/fosterTransport/pages"
	"net/http"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type")

	routeByMethod(w, r)
}

func routeByMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		routeGetByContentType(w, r)
	case "POST":
		routePostByContentType(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func routeGetByContentType(w http.ResponseWriter, r *http.Request) {
	pageContentTypes := []string{
		"text/html",
	}

	requestedContentTypes := r.Header.Get("Accept")

	for i := 0; i <= len(requestedContentTypes); i++ {
		if strings.Contains(requestedContentTypes, pageContentTypes[i]) {
			pages.Home(w, r)
			return
		}
	}

	http.Error(w, "Media type not supported", http.StatusUnsupportedMediaType)
}

func routePostByContentType(w http.ResponseWriter, r *http.Request) {
	pageContentTypes := []string{
		"application/x-www-form-urlencoded",
	}

	requestedContentTypes := r.Header.Get("Accept")

	for i := 0; i <= len(requestedContentTypes); i++ {
		if strings.Contains(requestedContentTypes, pageContentTypes[i]) {
			pages.Home(w, r)
			return
		}
	}

	http.Error(w, "Media type not supported", http.StatusUnsupportedMediaType)
}
