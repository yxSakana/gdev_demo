package model

type NovelQueryInput struct {
	Id        *uint64 `json:"id" form:"id"`
	Tag       *string `json:"tag" form:"tag"`
	Author    *string `json:"author" form:"author"`
	Title     *string `json:"title" form:"title"`
	WordCount *uint   `json:"word_count" form:"word_count"`
	View      *uint   `json:"view" form:"view"`
	Like      *uint   `json:"like" form:"like"`
}

type NovelOutput struct {
	NovelID      uint64   `json:"novel_id"`
	UserID       uint64   `json:"user_id"`
	Author       string   `json:"author"`
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	Description  string   `json:"description"`
	CoverUrl     string   `json:"cover_url"`
	Status       uint8    `json:"status"`
	WordCount    uint     `json:"word_count"`
	View         uint     `json:"view"`
	Like         uint     `json:"like"`
	ChapterCount uint     `json:"chapter_count"`
	ChapterIds   []uint64 `json:"chapter_ids"`
}
