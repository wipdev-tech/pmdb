// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

type Review struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	UserID       string
	MovieTmdbID  string
	Rating       int32
	Review       string
	PublicReview int32
}

type User struct {
	ID          string
	UserName    string
	DisplayName string
	Password    string
	Bio         string
}
