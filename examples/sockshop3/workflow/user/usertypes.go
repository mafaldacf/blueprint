package user

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type User struct {
	FirstName string    `json:"firstName" bson:"firstName"`
	LastName  string    `json:"lastName" bson:"lastName"`
	Email     string    `json:"email" bson:"email"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password,omitempty"`
	Addresses []Address `json:"addresses,omitempty" bson:"-"`
	Cards     []Card    `json:"cards,omitempty" bson:"-"`
	UserID    string    `json:"id" bson:"-"`
	Salt      string    `json:"salt" bson:"salt"`
}
type Address struct {
	Street   string
	Number   string
	Country  string
	City     string
	PostCode string
	ID       string
}
type Card struct {
	LongNum string
	Expires string
	CCV     string
	ID      string
}

func newUser() User {
	u := User{Addresses: make([]Address, 0), Cards: make([]Card, 0)}
	u.newSalt()
	return u
}

func (u *User) newSalt() {
	h := sha1.New()
	io.WriteString(h, strconv.Itoa(int(time.Now().UnixNano())))
	u.Salt = fmt.Sprintf("%x", h.Sum(nil))
}

func (u *User) cardIDs() []string {
	ids := []string{}
	for _, card := range u.Cards {
		ids = append(ids, card.ID)
	}
	return ids
}

func (u *User) addressIDs() []string {
	ids := []string{}
	for _, address := range u.Addresses {
		ids = append(ids, address.ID)
	}
	return ids
}

// Replace all CC numbers with asterisks, for returning to the user for display
func (u *User) maskCCs() {
	for i := range u.Cards {
		u.Cards[i].maskCC()
	}
}

// Replaces the CC number with asterisks for returning to the user for display
func (c *Card) maskCC() {
	l := len(c.LongNum) - 4
	c.LongNum = fmt.Sprintf("%v%v", strings.Repeat("*", l), c.LongNum[l:])
}
