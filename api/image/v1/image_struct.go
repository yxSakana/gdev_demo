package v1

import "time"

type ImageCollection struct {
	ID           uint64    `json:"id"`
	UploaderID   uint64    `json:"uploader_id"`
	UploaderName string    `json:"uploader_name"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CoverUrl     string    `json:"cover_url"`
	Number       int       `json:"number"`
	Tags         []string  `json:"tags"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
