package content

import (
	"fmt"
	"net/http"
	"fostertransport/datalayer"
)

func GetCss(w http.ResponseWriter, r *http.Request) {
	SetCacheHeaders(w, r, "text/css")
	
	sections, err := datalayer.GetSections(r)
	
	if err == nil {
		minifiedCss, err := MergeAndMinify("text/css", sections)
	
		if err == nil {
			fmt.Fprintf(w, minifiedCss)
			return
		}
	}
	
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func GetJavascript(w http.ResponseWriter, r *http.Request) {
	SetCacheHeaders(w, r, "text/javascript")
	
	sections, err := datalayer.GetSections(r)
	
	if err == nil {
		minifiedCss, err := MergeAndMinify("text/javascript", sections)
	
		if err == nil {
			fmt.Fprintf(w, minifiedCss)
			return
		}
	}
	
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
