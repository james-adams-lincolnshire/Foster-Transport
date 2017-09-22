package content

import (
	"net/http"
	"fostertransport/datalayer"
	"fostertransport/domain"
	"html/template"
	"log"
)

func GetCss(w http.ResponseWriter, r *http.Request) {
	SetCacheHeaders(w, r, "text/css")
	
	sections, err := datalayer.GetSections(r)
	
	if err == nil {
		minifiedCss, err := MergeAndMinify("text/css", sections)
		log.Println("log: " + minifiedCss + " ...")
		templateCss := template.HTML(minifiedCss)
		
		if err == nil {
			pageModel := domain.AdminPage {
				Name:  "content-layout",
				Model: domain.PageModel {
					Model: templateCss,
					Menu: domain.Menu {
						CurrentLocation: "content-layout",
						Sections: sections,
					},
				},
			}

			loadContentTemplate(w, pageModel)
			return
		}
	}
	
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func GetJavascript(w http.ResponseWriter, r *http.Request) {
	SetCacheHeaders(w, r, "text/javascript")
	
	sections, err := datalayer.GetSections(r)
	
	if err == nil {
		minifiedJavascript, err := MergeAndMinify("text/javascript", sections)
		
		templateJavascript := template.HTML(minifiedJavascript)
		
		if err == nil {
			pageModel := domain.AdminPage {
				Name:  "content-layout",
				Model: domain.PageModel {
					Model: templateJavascript,
					Menu: domain.Menu {
						CurrentLocation: "content-layout",
						Sections: sections,
					},
				},
			}

			loadContentTemplate(w, pageModel)
			return
		}
	}
	
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
