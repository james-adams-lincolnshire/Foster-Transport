package index

import (
	"fostertransport/datalayer"
	"fostertransport/domain"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	// Authenticate and authorize	
	if _, err := login(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	sections, err := datalayer.GetSections(r)
	if  err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	minifiedSections, err := MergeAndMinifyHtml(sections)
	if  err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	pageModel := domain.Page{
		Name:  "root",
		Model: minifiedSections,
	}

	loadTemplate(w, pageModel)
}
