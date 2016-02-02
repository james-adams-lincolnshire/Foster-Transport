package domain

import (
	"appengine/datastore"
)

type Return struct {
	Key  *datastore.Key
	Id   int64
	Data interface{}
}
