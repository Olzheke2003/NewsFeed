package models

import "time"

type Users struct {
	ID                int       `json:id`
	Email             string    `json:email`
	Username          string    `json:username`
	Password          string    `json:password`
	EncryptedPassword string    `json:"-"`
	Image             string    `json:"image"`
	Name              string    `json:name`
	Lastname          string    `json:lastname`
	DateOfBirth       string    `json:dateofbirth`
	Gender            string    `json:gender`
	CreatedAt         time.Time `json:create_at`
}
