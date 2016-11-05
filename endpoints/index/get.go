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
	
	// Get visible sections
	sections, err := datalayer.GetSections(r)
	if  err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	for i := len(sections) - 1; i >= 0; i-- {
		section := sections[i]
		
		if section.Hidden {
			sections = append(sections[:i], sections[i+1:]...)
		}
	}
	
	// Optimize HTML
	minifiedSections, err := MergeAndMinifyHtml(sections)
	if  err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Render
	pageModel := domain.Page{
		Name:  "root",
		Model: minifiedSections,
	}

	loadTemplate(w, pageModel)
}
