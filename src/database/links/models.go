// Code generated by sqlc. DO NOT EDIT.

package links

import (
	"database/sql"
)

type Link struct {
	ID     int64  `json:"id"`
	Code   string `json:"code"`
	UserID int64  `json:"user_id"`
}

type User struct {
	ID           int64         `json:"id"`
	Firstname    string        `json:"firstname"`
	Lastname     string        `json:"lastname"`
	Email        string        `json:"email"`
	Upassword    string        `json:"upassword"`
	Isambassador sql.NullInt32 `json:"isambassador"`
}
