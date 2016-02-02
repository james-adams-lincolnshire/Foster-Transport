package datalayer

import (
	"appengine"
	"appengine/datastore"
	"fostertransport/domain"
	"log"
	"net/http"
)

func GetPages(w http.ResponseWriter, r *http.Request) []domain.Return {
	context := appengine.NewContext(r)

	results := make([]domain.Page, 0, 10)

	query := datastore.NewQuery("Pages").Order("-Timestamp")
	log.Println("log: Executing datastore query for Pages...")
	keys, err := query.GetAll(context, &results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	res := make([]domain.Return, 0, 10)
	for i, r := range results {
		k := keys[i]
		y := domain.Return{
			Key:  k,
			Id:   k.IntID(),
			Data: r,
		}
		res = append(res, y)
	}

	return res
}
