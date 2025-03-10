package entity

import "time"

type ImageCollection struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint64    `gorm:"not null" json:"user_id"`
	Uploader    string    `gorm:"size:50;not null" json:"uploader"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `json:"description"`
	CoverUrl    string    `gorm:"size:255" json:"cover_url"`
	Number      int       `gorm:"not null" json:"number"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Image struct {
	ID           uint64 `gorm:"primary_key;auto_increment" json:"id"`
	CollectionID uint64 `gorm:"not null" json:"collection_id"`
	ImageUrl     string `gorm:"size:255;not null" json:"image_url"`
}

type ImageTag struct {
	ID   uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:50;not null" json:"name"`
}

type ImageTagRel struct {
	CollectionID uint64 `gorm:"primary_key" json:"collection_id"`
	ImageTagID   uint64 `gorm:"primary_key" json:"image_tag_id"`
}

func (ImageTagRel) TableName() string {
	return "image_tag_rel"
}
