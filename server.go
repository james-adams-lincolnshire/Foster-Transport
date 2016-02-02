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

	router.HandleFunc("/content/css/site.css", content.GetCss).Methods("GET")
	router.HandleFunc("/content/javascript/site.js", content.GetJavascript).Methods("GET")

	router.HandleFunc("/admin/dashboard", admin.GetDashboard).Methods("GET")
	router.HandleFunc("/admin/quote/chats", admin.GetLiveChats).Methods("GET")
	router.HandleFunc("/admin/quote/requests", admin.GetQuoteRequests).Methods("GET")
	router.HandleFunc("/admin/manage/sections", admin.GetManageSections).Methods("GET")
	router.HandleFunc("/admin/manage/sections/create", admin.GetCreateSection).Methods("GET")
	router.HandleFunc("/admin/manage/sections/edit", admin.GetEditSection).Methods("GET")
	router.HandleFunc("/admin/manage/blog", admin.GetManageBlog).Methods("GET")
	router.HandleFunc("/admin/manage/blog/create", admin.GetCreateBlog).Methods("GET")

	router.HandleFunc("/admin/manage/sections", admin.PostSaveSection).Methods("POST")
	router.HandleFunc("/admin/manage/blog", admin.PostSaveBlog).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("log: HTTP server listening...")
	http.Handle("/", router)
}
