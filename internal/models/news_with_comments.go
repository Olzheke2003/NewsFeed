package models

type NewsWithComments struct {
	News          News `json:"news"`
	CommentsCount int  `json:"comments_count"`
}
