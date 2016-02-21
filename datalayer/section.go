package datalayer

import (
	"appengine"
	"appengine/datastore"
	"fostertransport/domain"
	"log"
	"net/http"
)

func GetSections(r *http.Request) ([]domain.Section, error) {
	context := appengine.NewContext(r)
	sections := make([]domain.Section, 0, 100)
	query := datastore.NewQuery("Section").Order("Order")

	log.Println("log: Executing datastore query for Sections...")
	keys, err := query.GetAll(context, &sections)

	for index, _ := range sections {
		key := keys[index]

		sections[index].Key = key
		sections[index].Id = key.IntID()
	}

	return sections, err
}

func GetSection(r *http.Request, id int64) (domain.Section, error) {
	context := appengine.NewContext(r)
	section := domain.Section{}
	key := sectionKey(context, id)
	log.Println("log: Executing datastore get for a Section...")
	err := datastore.Get(context, key, &section)

	if err == nil {
		section.Key = key
		section.Id = key.IntID()
	}

	return section, err
}

func CreateSection(r *http.Request, section *domain.Section) error {
	context := appengine.NewContext(r)

	log.Println("log: Executing datastore query to create a Section...")
	incKey := datastore.NewIncompleteKey(context, "Section", nil)
	key, err := datastore.Put(context, incKey, section)

	if err == nil {
		section.Key = key
		section.Id = key.IntID()
	}

	return err
}

func EditSection(r *http.Request, id int64, section *domain.Section) error {
	context := appengine.NewContext(r)

	log.Println("log: Executing datastore query to edit a Section...")
	key := sectionKey(context, id)
	_, err := datastore.Put(context, key, section)

	return err
}

func DeleteSection(r *http.Request, id int64) error {
	context := appengine.NewContext(r)

	log.Println("log: Executing datastore query to delete a Section...")
	key := sectionKey(context, id)
	err := datastore.Delete(context, key)

	return err
}

func newSectionKey(c appengine.Context) *datastore.Key {
	// The string "default_sections" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Section", "", 0, nil)
}

func sectionKey(c appengine.Context, id int64) *datastore.Key {
	return datastore.NewKey(c, "Section", "", id, nil)
}
