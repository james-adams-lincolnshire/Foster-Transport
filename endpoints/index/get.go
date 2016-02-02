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
	
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.Page{
			Name:  "root",
			Model: sections,
		}

		loadTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
