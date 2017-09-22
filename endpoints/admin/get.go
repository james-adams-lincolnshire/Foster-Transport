package admin

import (
	"fostertransport/datalayer"
	"fostertransport/domain"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {		
		pageModel := domain.AdminPage{
			Name:  "dashboard",
			Model: domain.PageModel{
				Model: sections,
				Menu: domain.Menu{
					CurrentLocation: "dashboard",
					Sections: sections,
				},
			},
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetManageSections(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		if len(sections) == 0 {
			http.Redirect(w, r, "/admin/sections/create", 307)
			return
		}
	
		head := make([]struct {
			Id int64
			Html template.HTML
			Css template.CSS
		}, 1, 1)
		htmlSections := make([]struct {
			Id int64
			Html template.HTML
			Css template.CSS
		}, len(sections) - 2, len(sections) - 2)
		deferredContent := make([]struct {
			Id int64
			Html template.HTML
			Css template.CSS
		}, 1, 1)
		
		htmlSectionsIndex := 0
		
		for i := 0; i < len(sections); i++ {
			section := sections[i]
			
			if section.Name == "Head" {
				head[0].Id = section.Id
				head[0].Html = template.HTML(section.Html)
				head[0].Css = template.CSS(section.Css)
			} else if  section.Name == "Deferred Content" {
				deferredContent[0].Id = section.Id
				deferredContent[0].Html = template.HTML(section.Html)
				deferredContent[0].Css = template.CSS(section.Css)
			} else {
				htmlSections[htmlSectionsIndex].Id = section.Id
				htmlSections[htmlSectionsIndex].Html = template.HTML(section.Html)
				htmlSections[htmlSectionsIndex].Css = template.CSS(section.Css)
				htmlSectionsIndex++
			}
		}
		
		model := make(map[string][]struct {
			Id int64
			Html template.HTML
			Css template.CSS
		})
	
		model["Head"] = head
		model["Sections"] = htmlSections
		model["DeferredContent"] = deferredContent
		
		pageModel := domain.AdminPage{
			Name:  "manage-sections",
			Model: domain.PageModel{
				Model: model,
				Menu: domain.Menu{
					CurrentLocation: "manage-sections",
					Sections: sections,
				},
			},
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCreateSection(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.AdminPage {
			Name:  "create-section",
			Model: domain.PageModel {
				Model: sections,
				Menu: domain.Menu {
					CurrentLocation: "create-section",
					Sections: sections,
				},
			},
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
	sections, err := datalayer.GetSections(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	pageModel := domain.AdminPage {
		Name:  "edit-section",
		Model: domain.PageModel {
			Model: section,
			Menu: domain.Menu {
				CurrentLocation: "edit-section",
				Sections: sections,
			},
		},
	}

	loadAdminTemplate(w, pageModel)
}

func GetManageBlog(w http.ResponseWriter, r *http.Request) {
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

func GetCreateBlog(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.AdminPage {
			Name:  "create-blog",
			Model: domain.PageModel {
				Model: sections,
				Menu: domain.Menu {
					CurrentLocation: "create-blog",
					Sections: sections,
				},
			},
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetLiveChats(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.AdminPage {
			Name:  "live-chats",
			Model: domain.PageModel {
				Model: sections,
				Menu: domain.Menu {
					CurrentLocation: "live-chats",
					Sections: sections,
				},
			},
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetQuoteRequests(w http.ResponseWriter, r *http.Request) {
	if sections, err := datalayer.GetSections(r); err == nil {
		pageModel := domain.AdminPage {
			Name:  "quote-requests",
			Model: domain.PageModel {
				Model: sections,
				Menu: domain.Menu {
					CurrentLocation: "quote-requests",
					Sections: sections,
				},
			},
		}

		loadAdminTemplate(w, pageModel)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
