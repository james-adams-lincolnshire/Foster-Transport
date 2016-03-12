package content

import (
	"fmt"
	"net/http"
	"time"
	"fostertransport/domain"
	"bytes"
	"html/template"
	"log"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
    "github.com/tdewolff/minify/js"
)

var contentTemplates = make(map[string]*template.Template)

func SetCacheHeaders(w http.ResponseWriter, r *http.Request, contentType string) {
	dateTime := time.Now()
	cacheExpiry := dateTime.Add((time.Hour * 24))
	cacheExpiryMil := fmt.Sprintf("%v", 30000)

	w.Header().Add("Content-Type", contentType)
	w.Header().Add("Expires", cacheExpiry.Format("Mon, 02 Jan 2006 15:04:05 MST"))
	w.Header().Set("Cache-Control", "public, max-age=86400 ,s-maxage="+cacheExpiryMil)
	w.Header().Set("Pragma", "Public")
}

func MergeAndMinify(contentType string, sections []domain.Section) (string, error) {
	var buffer bytes.Buffer
	minifier := minify.New()
	
	if contentType == "text/css" {
		minifier.AddFunc(contentType, css.Minify)
	
		for i := 0; i < len(sections); i++ {
			buffer.WriteString(sections[i].Css)
		}
	}
	
	if contentType == "text/javascript" {
		minifier.AddFunc(contentType, js.Minify)
	
		for i := 0; i < len(sections); i++ {
			buffer.WriteString(sections[i].Javascript)
		}
	}
	
	return minifier.String(contentType, buffer.String())
}

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

func loadContentTemplate(w http.ResponseWriter, adminPage domain.AdminPage) {
	log.Println("log: Executing template, content...")

	cacheContentTemplates()

	err := contentTemplates[adminPage.Name].ExecuteTemplate(w, "layout", adminPage.Model)
	
	if err != nil {
		log.Println("error? " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func cacheContentTemplates() {
	if len(contentTemplates) > 0 {
		return
	}

	contentTemplates["content-layout"] = template.Must(template.ParseFiles(
		"views/admin/contentLayout.html",
	))
}