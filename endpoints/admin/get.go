package admin

import (
	"fostertransport/datalayer"
	"fostertransport/domain"
	"log"
	"net/http"
	"strconv"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.AdminPage{
			Name:  "dashboard",
			Model: sections,
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetManageSections(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		head := make([]domain.Section, 1, 1)
		
		for i := len(sections) - 1; i >= 0; i-- {
			section := sections[i]
			
			if section.Name == "Head" {
				head[0] = section
				sections = append(sections[:i], sections[i+1:]...)
				break
			}
		}
		
		model := make(map[string][]domain.Section)
	
		model["Head"] = head
		model["Sections"] = sections
	
		pageModel := domain.AdminPage{
			Name:	"manage-sections",
			Model:	model,
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCreateSection(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.AdminPage{
			Name:  "create-section",
			Model: sections,
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetEditSection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)

	log.Println("log: Provided id " + r.FormValue("id") + " ...")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	section, err := datalayer.GetSection(r, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageModel := domain.AdminPage{
		Name:  "edit-section",
		Model: section,
	}

	loadAdminTemplate(w, pageModel)
}

func GetManageBlog(w http.ResponseWriter, r *http.Request) {
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

func GetCreateBlog(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.AdminPage{
			Name:  "create-blog",
			Model: sections,
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetLiveChats(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.AdminPage{
			Name:  "live-chats",
			Model: sections,
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetQuoteRequests(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.AdminPage{
			Name:  "quote-requests",
			Model: sections,
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
