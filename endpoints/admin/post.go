package admin

import (
	"fostertransport/datalayer"
	"fostertransport/domain"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func PostSaveSection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	editing := r.URL.Query().Get("id") != ""
	
	if err != nil && editing {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order, _ := strconv.Atoi(r.FormValue("displayOrder"))
	hidden := convertCheckbox(r.FormValue("hidden"))
	deleted := convertCheckbox(r.FormValue("deleted"))
	aboveTheFold := convertCheckbox(r.FormValue("aboveTheFold"))

	var section = domain.Section{
		Order:   -1,
		Hidden:  true,
		AboveTheFold: false,
		Created: time.Now(),
	}

	if editing {
		if (deleted) {
			if err := datalayer.DeleteSection(r, id); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			
			http.Redirect(w, r, "/admin/sections", 302)
			return
		}
	
		editableSection, err := datalayer.GetSection(r, id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		section = editableSection
		
		if _, exists := r.Form["displayOrder"]; exists { section.Order = order }
		section.Hidden = hidden
		section.AboveTheFold = aboveTheFold
	}

	if _, exists := r.Form["sectionName"]; exists { section.Name = r.FormValue("sectionName") }
	if _, exists := r.Form["html"]; exists { 
		if htmlVal, err := url.QueryUnescape(r.FormValue("html")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			section.Html = strings.TrimSpace(htmlVal)
		}
	}
	if _, exists := r.Form["css"]; exists { section.Css = r.FormValue("css") }
	if _, exists := r.Form["javascript"]; exists { section.Javascript = r.FormValue("javascript")}

	if editing {
		if err := datalayer.EditSection(r, id, &section); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := datalayer.CreateSection(r, &section); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/admin/sections", 302)
}

func PostSaveBlog(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {		
		pageModel := domain.AdminPage {
			Name:  "manage-blog",
			Model: domain.PageModel {
				Model: sections,
				Menu: domain.Menu {
					CurrentLocation: "manage-blog",
					Sections: sections,
				},
			},
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
