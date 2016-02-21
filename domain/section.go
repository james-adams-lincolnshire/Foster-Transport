package domain

import (
	"appengine/datastore"
	"html/template"
	"time"
)

type Section struct {
	Key			*datastore.Key	`datastore:"-"`
	Id				int64					`datastore:"-"`
	Name		string
	Order		int
	Hidden		bool
	Created	time.Time
	Html			template.HTML	`datastore:",noindex"`
	Css			string				`datastore:",noindex"`
	Javascript	string				`datastore:",noindex"`
}
