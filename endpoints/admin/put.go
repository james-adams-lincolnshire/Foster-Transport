package admin

import (
	"fostertransport/datalayer"
	"html/template"
	"net/http"
	"strconv"
)

func PutEditSection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	order, err := strconv.Atoi(r.FormValue("order"))
	hidden := convertCheckbox(r.FormValue("hidden"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	section, err := datalayer.GetSection(r, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	section.Name = r.FormValue("sectionName")
	section.Order = order
	section.Hidden = hidden
	section.Html = template.HTML(r.FormValue("html"))
	section.Css = r.FormValue("css")
	section.Javascript = r.FormValue("javascript")

	if err := datalayer.EditSection(r, id, &section); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/admin/manage/sections", 302)
	}
}
