package index

import (
	"fostertransport/domain"
	"fostertransport/datalayer"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/securecookie"
)

var templates = make(map[string]*template.Template)
var hashKey = []byte(securecookie.GenerateRandomKey(64))
var blockKey = []byte(securecookie.GenerateRandomKey(32))

func loadAdhoctemplate(w http.ResponseWriter, adhocTemplate string, sections []domain.Section) {
	log.Println("log: Executing template, adhoc...")

	parsedTemplate, err := template.New("adhoc").Parse(adhocTemplate)

	if err == nil {
		if err = parsedTemplate.Execute(w, sections); err == nil {
			return
		}
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func loadTemplate(w http.ResponseWriter, page domain.Page) {
	log.Println("log: Executing template...")

	cacheTemplates()

	if err := templates[page.Name].ExecuteTemplate(w, "layout", page.Model); err != nil {
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

func cacheTemplates() {
	if len(templates) > 0 {
		return
	}

	templates["root"] = template.Must(template.ParseFiles(
		"views/root.html",
	))
}

func login(w http.ResponseWriter, r *http.Request) (domain.User, error) {
	thisUser := domain.NewUser()
	cookieName := "ft-auth"
	
	cookie, err := getCookie(w, r, cookieName)
	
	if err == nil {
		thisUser, err = datalayer.GetUser(r, cookie)
	} else {
		if err = datalayer.CreateUser(r, &thisUser); err == nil {
			err = setCookie(w, r, thisUser.Cookies[0], cookieName)
		}
	}
	
	if err == nil {
		thisUser.Login()
		err = datalayer.EditUser(r, thisUser.Id, &thisUser)
	}
	
	return thisUser, err
}

func setCookie(w http.ResponseWriter, r *http.Request, value string, cookieName string) error {
	secureCookie := securecookie.New(hashKey, blockKey)
	encoded, err := secureCookie.Encode(cookieName, value)
	
    if err == nil {
        cookie := &http.Cookie{
            Name:  cookieName,
            Value: encoded,
            Path:  "/",
			HttpOnly: true,
			//Secure: true,
			MaxAge: 31540000,
        }
		
        http.SetCookie(w, cookie)
    }
	
	return err
}

func getCookie(w http.ResponseWriter, r *http.Request, cookieName string) (string, error) {
	value := ""
	cookie, err := r.Cookie(cookieName)
	
    if err == nil {
        secureCookie := securecookie.New(hashKey, blockKey)
        err = secureCookie.Decode(cookieName, cookie.Value, &value)
    }
	
	return value, err
}