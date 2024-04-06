// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: reviews.sql

package database

import (
	"context"
)

const createReview = `-- name: CreateReview :one
INSERT INTO reviews (
    id,
    created_at,
    updated_at,
    user_id,
    movie_tmdb_id,
    rating,
    review,
    public_review
)
VALUES ( ?, ?, ?, ?, ?, ?, ?, ? )
RETURNING id, created_at, updated_at, user_id, movie_tmdb_id, rating, review, public_review
`

type CreateReviewParams struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	UserID       string
	MovieTmdbID  string
	Rating       int64
	Review       string
	PublicReview int64
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error) {
	row := q.db.QueryRowContext(ctx, createReview,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.MovieTmdbID,
		arg.Rating,
		arg.Review,
		arg.PublicReview,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.MovieTmdbID,
		&i.Rating,
		&i.Review,
		&i.PublicReview,
	)
	return i, err
}

const getReviews = `-- name: GetReviews :many
SELECT
    r.id,
    r.user_id,
    u.display_name as user_name,
    r.created_at,
    r.updated_at,
    r.movie_tmdb_id,
    r.rating,
    r.review 
FROM reviews r
JOIN users u
ON u.id = r.user_id
WHERE public_review = 1
ORDER BY r.created_at DESC
LIMIT 5
`

type GetReviewsRow struct {
	ID          string
	UserID      string
	UserName    string
	CreatedAt   string
	UpdatedAt   string
	MovieTmdbID string
	Rating      int64
	Review      string
}

func (q *Queries) GetReviews(ctx context.Context) ([]GetReviewsRow, error) {
	rows, err := q.db.QueryContext(ctx, getReviews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetReviewsRow
	for rows.Next() {
		var i GetReviewsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.UserName,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MovieTmdbID,
			&i.Rating,
			&i.Review,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReviewsForMovie = `-- name: GetReviewsForMovie :many
SELECT
    r.id,
    r.user_id,
    u.display_name as user_name,
    r.created_at,
    r.updated_at,
    r.movie_tmdb_id,
    r.rating,
    r.review 
FROM reviews r
JOIN users u
ON u.id = r.user_id
WHERE movie_tmdb_id = ? AND public_review = 1
`

type GetReviewsForMovieRow struct {
	ID          string
	UserID      string
	UserName    string
	CreatedAt   string
	UpdatedAt   string
	MovieTmdbID string
	Rating      int64
	Review      string
}

func (q *Queries) GetReviewsForMovie(ctx context.Context, movieTmdbID string) ([]GetReviewsForMovieRow, error) {
	rows, err := q.db.QueryContext(ctx, getReviewsForMovie, movieTmdbID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetReviewsForMovieRow
	for rows.Next() {
		var i GetReviewsForMovieRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.UserName,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MovieTmdbID,
			&i.Rating,
			&i.Review,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
