package datalayer

import (
	"appengine"
	"appengine/datastore"
	"fostertransport/domain"
	"log"
	"net/http"
)

func GetUsers(r *http.Request) ([]domain.User, error) {
	context := appengine.NewContext(r)
	users := make([]domain.User, 0, 10)
	query := datastore.NewQuery("User").Order("Created")

	log.Println("log: Executing datastore query for Users...")
	keys, err := query.GetAll(context, &users)

	for index, _ := range users {
		key := keys[index]

		users[index].Key = key
		users[index].Id = key.IntID()
	}

	return users, err
}

func GetUser(r *http.Request, cookie string) (domain.User, error) {
	context := appengine.NewContext(r)
	users := make([]domain.User, 0, 1)
	query := datastore.NewQuery("User").Filter("Cookies =", cookie).Order("Created")

	log.Println("log: Executing datastore query for a User...")
	keys, err := query.GetAll(context, &users)

	if err == nil {
		for index, _ := range users {
			key := keys[index]

			users[index].Key = key
			users[index].Id = key.IntID()
		}
	}

	return users[0], err
}

func CreateUser(r *http.Request, thisUser *domain.User) error {
	context := appengine.NewContext(r)

	log.Println("log: Executing datastore query to create a User...")
	incKey := datastore.NewIncompleteKey(context, "User", nil)
	key, err := datastore.Put(context, incKey, thisUser)

	if err == nil {
		thisUser.Key = key
		thisUser.Id = key.IntID()
	}

	return err
}

func EditUser(r *http.Request, id int64, thisUser *domain.User) error {
	context := appengine.NewContext(r)

	log.Println("log: Executing datastore query to edit a User...")
	key := userKey(context, id)
	_, err := datastore.Put(context, key, thisUser)

	return err
}

func DeleteUser(r *http.Request, id int64) error {
	context := appengine.NewContext(r)

	log.Println("log: Executing datastore query to delete a User...")
	key := userKey(context, id)
	err := datastore.Delete(context, key)

	return err
}

func newUserKey(c appengine.Context) *datastore.Key {
	// The string "default_sections" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "User", "", 0, nil)
}

func userKey(c appengine.Context, id int64) *datastore.Key {
	return datastore.NewKey(c, "User", "", id, nil)
}