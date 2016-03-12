package admin

import (
	"fostertransport/datalayer"
	"fostertransport/domain"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func PostSaveSection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	editing := r.URL.Query().Get("id") != ""
	
	if err != nil && editing {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order, err := strconv.Atoi(r.FormValue("displayOrder"))
	hidden := convertCheckbox(r.FormValue("hidden"))
	deleted := convertCheckbox(r.FormValue("deleted"))

	if err != nil && editing {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var section = domain.Section{
		Order:   -1,
		Hidden:  true,
		Created: time.Now(),
	}

	if editing {
		if (deleted) {
			if err := datalayer.DeleteSection(r, id); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			
			http.Redirect(w, r, "/admin/manage/sections", 302)
			return
		}
	
		editableSection, err := datalayer.GetSection(r, id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		section = editableSection
		
		section.Order = order
		section.Hidden = hidden
	}

	section.Name = r.FormValue("sectionName")
	section.Html = template.HTML(r.FormValue("html"))
	section.Css = r.FormValue("css")
	section.Javascript = r.FormValue("javascript")

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

	http.Redirect(w, r, "/admin/manage/sections", 302)
}

func PostSaveBlog(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.AdminPage{
			Name:  "manage-blog",
			Model: sections,
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
