package entity

import "time"

type Novel struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	Uploader    string    `json:"Uploader"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CoverUrl    string    `json:"cover_url"`
	Status      uint8     `json:"status"`
	WordCount   uint      `json:"word_count"`
	View        uint      `json:"view"`
	Like        uint      `json:"like"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NovelChapter struct {
	ID        uint64    `json:"id"`
	NovelId   string    `json:"novel_id"`
	Title     string    `json:"title"`
	Number    int       `json:"num"`
	Content   string    `json:"content"`
	WordCount uint      `json:"word_count"`
	View      uint      `json:"view"`
	Like      uint      `json:"like"`
	CreatedAt time.Time `json:"created_at"`
}

type NovelTag struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type NovelTagRel struct {
	NovelId    uint64 `gorm:"primaryKey" json:"novel_id"`
	NovelTagId uint64 `gorm:"primaryKey" json:"novel_tag_id"`
}

func (NovelTagRel) TableName() string {
	return "novel_tag_rel"
}
