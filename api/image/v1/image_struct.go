package v1

import "time"

type ImageCollection struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"uploader_id"`
	Uploader    string    `json:"uploader_name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CoverUrl    string    `json:"cover_url"`
	Tags        []string  `json:"tags"`
	Number      int       `json:"number"`
	View        int       `json:"view" redis:"view"`
	Like        int       `json:"like" redis:"like"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
