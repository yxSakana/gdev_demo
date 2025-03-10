package conv

import (
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
)

func UserToEntity(in do.User) entity.User {
	return *in.User
}

func UserToDo(in entity.User) do.User {
	return do.User{User: &in}
}

//func ToEntity(in do.User) entity.User {
//h := func(v string) *string {
//	if v != "" {
//		return &v
//	}
//	return nil
//}
//return entity.User{
//	ID:         in.ID,
//	Username:   in.Username,
//	Nickname:   in.Nickname,
//	Password:   in.Password,
//	Email:      h(in.Email),
//	PictureUrl: h(in.PictureUrl),
//	Role:       in.Role,
//	CreatedAt:  in.CreatedAt,
//}
//}

//func ToDo(in entity.User) do.User {
//h := func(ptr *string) string {
//	if ptr != nil {
//		return *ptr
//	}
//	return ""
//}
//return do.User{
//	ID:         in.ID,
//	Username:   in.Username,
//	Nickname:   in.Nickname,
//	Password:   in.Password,
//	Email:      h(in.Email),
//	PictureUrl: h(in.PictureUrl),
//	Role:       in.Role,
//	CreatedAt:  in.CreatedAt,
//}
//}
