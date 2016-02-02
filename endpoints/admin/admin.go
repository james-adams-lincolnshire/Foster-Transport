package admin

import (
	"fostertransport/domain"
	"html/template"
	"log"
	"net/http"
)

var adminTemplates = make(map[string]*template.Template)

func loadAdhoctemplate(w http.ResponseWriter, adhocTemplate string, sections []domain.Section) {
	log.Println("log: Executing template, adhoc...")

	parsedTemplate, err := template.New("adhoc").Parse(adhocTemplate)

	if err == nil {
		if err := parsedTemplate.Execute(w, sections); err == nil {
			return
		}
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func loadAdminTemplate(w http.ResponseWriter, adminPage domain.AdminPage) {
	log.Println("log: Executing template, admin...")

	cacheAdminTemplates()

	if err := adminTemplates[adminPage.Name].ExecuteTemplate(w, "layout", adminPage.Model); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func convertCheckbox(value string) bool {
	if value == "on" {
		return true
	}

	return false
}

func cacheAdminTemplates() {
	if len(adminTemplates) > 0 {
		return
	}

	adminTemplates["dashboard"] = template.Must(template.ParseFiles(
		"views/admin/dashboard.html",
		"views/admin/header.html",
		"views/admin/head.html",
		"views/admin/layout.html",
	))

	adminTemplates["live-chats"] = template.Must(template.ParseFiles(
		"views/admin/liveChats.html",
		"views/admin/header.html",
		"views/admin/head.html",
		"views/admin/layout.html",
	))

	adminTemplates["quote-requests"] = template.Must(template.ParseFiles(
		"views/admin/quoteRequests.html",
		"views/admin/header.html",
		"views/admin/head.html",
		"views/admin/layout.html",
	))

	adminTemplates["manage-sections"] = template.Must(template.ParseFiles(
		"views/admin/manageSections.html",
		"views/admin/header.html",
		"views/admin/head.html",
		"views/admin/layout.html",
	))

	adminTemplates["create-section"] = template.Must(template.ParseFiles(
		"views/admin/createSection.html",
		"views/admin/header.html",
		"views/admin/head.html",
		"views/admin/layout.html",
	))

	adminTemplates["edit-section"] = template.Must(template.ParseFiles(
		"views/admin/editSection.html",
		"views/admin/header.html",
		"views/admin/head.html",
		"views/admin/layout.html",
	))

	adminTemplates["manage-blog"] = template.Must(template.ParseFiles(
		"views/admin/manageBlog.html",
		"views/admin/header.html",
		"views/admin/head.html",
		"views/admin/layout.html",
	))

	adminTemplates["create-blog"] = template.Must(template.ParseFiles(
		"views/admin/createBlog.html",
		"views/admin/header.html",
		"views/admin/head.html",
		"views/admin/layout.html",
	))
}
