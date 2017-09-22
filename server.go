package fostertransport

import (
	"fostertransport/endpoints/admin"
	"fostertransport/endpoints/content"
	"fostertransport/endpoints/index"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/", index.GetIndex).Methods("GET")
	router.HandleFunc("/our-services", index.GetIndex).Methods("GET")

	router.HandleFunc("/content/css/site.css", content.GetCss).Methods("GET")
	router.HandleFunc("/content/javascript/site.js", content.GetJavascript).Methods("GET")

	router.HandleFunc("/admin/dashboard", admin.GetDashboard).Methods("GET")
	router.HandleFunc("/admin/chats", admin.GetLiveChats).Methods("GET")
	router.HandleFunc("/admin/quotes", admin.GetQuoteRequests).Methods("GET")
	router.HandleFunc("/admin/sections", admin.GetManageSections).Methods("GET")
	router.HandleFunc("/admin/sections/create", admin.GetCreateSection).Methods("GET")
	router.HandleFunc("/admin/sections/edit", admin.GetEditSection).Methods("GET")
	router.HandleFunc("/admin/blog", admin.GetManageBlog).Methods("GET")
	router.HandleFunc("/admin/blog/create", admin.GetCreateBlog).Methods("GET")

	router.HandleFunc("/admin/sections", admin.PostSaveSection).Methods("POST")
	router.HandleFunc("/admin/blog", admin.PostSaveBlog).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("log: HTTP server listening...")
	http.Handle("/", router)
}
