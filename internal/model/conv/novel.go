package conv

import (
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
)

func NovelToEntity(in do.Novel) entity.Novel {
	return entity.Novel{
		ID:            in.ID,
		UserID:        in.UserID,
		Uploader:      in.Uploader,
		Title:         in.Title,
		Description:   in.Description,
		CoverUrl:      in.CoverUrl,
		Status:        in.Status,
		ChapterNumber: in.ChapterNumber,
		WordCount:     in.WordCount,
		View:          in.View,
		Like:          in.Like,
		CreatedAt:     in.CreatedAt,
		UpdatedAt:     in.UpdatedAt,
	}
}

func NovelToDo(in entity.Novel, tags []string) do.Novel {
	return do.Novel{
		ID:            in.ID,
		UserID:        in.UserID,
		Uploader:      in.Uploader,
		Title:         in.Title,
		Description:   in.Description,
		CoverUrl:      in.CoverUrl,
		Tags:          tags,
		Status:        in.Status,
		ChapterNumber: in.ChapterNumber,
		WordCount:     in.WordCount,
		View:          in.View,
		Like:          in.Like,
		CreatedAt:     in.CreatedAt,
		UpdatedAt:     in.UpdatedAt,
	}
}

func NovelChapterToEntity(in do.NovelChapter) entity.NovelChapter {
	return entity.NovelChapter{
		ID:        in.ID,
		NovelId:   in.NovelId,
		Title:     in.Title,
		Number:    in.Number,
		Content:   in.Content,
		WordCount: in.WordCount,
		View:      in.View,
		Like:      in.Like,
		CreatedAt: in.CreatedAt,
	}
}
