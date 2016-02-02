package domain

import (
	"appengine/datastore"
	"time"
	"log"
)

type User struct {
	Key			*datastore.Key `datastore:"-"`
	Id				int64          `datastore:"-"`
	Cookies	[]string
	Logins		[]time.Time
	Created	time.Time
}

func NewUser() User {
	cookies := []string {
		NewUUID().String(),
	}
	
	return User{
		Cookies: cookies,
		Created: time.Now(),
	}
}

func (user *User) Login() {
	for i := 0; i < len(user.Logins); i++ {
		log.Println(user.Logins[i].Format(time.RFC3339))
	}

	user.Logins = append(user.Logins, time.Now())
	
	log.Println("log: User login")
	
	for i := 0; i < len(user.Logins); i++ {
		log.Println(user.Logins[i].Format(time.RFC3339))
	}
}
