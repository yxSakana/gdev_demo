package v1

import "time"

type ImageCollection struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"uploader_id"`
	Uploader    string    `json:"uploader_name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CoverUrl    string    `json:"cover_url"`
	Number      int       `json:"number"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
