package models

import "time"

type Gender struct {
	ID     int    `json:"id"`
	Gender string `json:gender`
}

type Users struct {
	ID                int       `json:id`
	Email             string    `json:email`
	Username          string    `json:username`
	Password          string    `json:password`
	EncryptedPassword string    `json:"-"`
	Image             string    `json:"image"`
	Name              string    `json:name`
	Lastname          string    `json:lastname`
	DateOfBirth       time.Time `json:dateofbirth`
	GenderID          int       `json:gender_id`
	CreatedAt         time.Time `json:create_at`
}
