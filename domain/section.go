package domain

import (
	"appengine/datastore"
	"time"
)

type Section struct {
	Key				*datastore.Key	`datastore:"-"`
	Id					int64					`datastore:"-"`
	Name			string
	Order			int
	Hidden			bool
	AboveTheFold bool
	Created		time.Time
	Html				string	`datastore:",noindex"`
	Css				string	`datastore:",noindex"`
	Javascript		string	`datastore:",noindex"`
}