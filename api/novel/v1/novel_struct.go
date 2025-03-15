package v1

type NovelDetail struct {
	NovelID       uint64   `json:"novel_id"`
	UploaderID    uint64   `json:"user_id"`
	Author        string   `json:"author"`
	Title         string   `json:"title"`
	Tags          []string `json:"tags"`
	Description   string   `json:"description"`
	CoverUrl      string   `json:"cover_url"`
	Status        uint8    `json:"status"`
	WordCount     uint     `json:"word_count"`
	View          uint     `json:"view"`
	Like          uint     `json:"like"`
	ChapterNumber uint     `json:"chapter_number"`
}
