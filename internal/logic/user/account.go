package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yxSakana/gdev_demo/internal/dao"
	"github.com/yxSakana/gdev_demo/internal/dao/user"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"github.com/yxSakana/gdev_demo/utility"
)

func Login(c *gin.Context, username, password string) (token string, err error) {
	userEntity, err := user.GetUserByUsername(dao.Ctx(c), username)
	if err != nil {
		return "", err
	}

	ok := utility.CheckPassword(password, userEntity.Password)
	if !ok {
		return "", errors.New("password error")
	}

	return GenerateToken(userEntity.ID)
}

func Register(c *gin.Context, u *do.User) error {
	u.Password = utility.MustGeneratePassword(u.Password)
	return user.Create(dao.Ctx(c), u)
}

func IsExist(c *gin.Context, uid uint64) (bool, error) {
	_, err := user.GetUserByID(dao.Ctx(c), uid)
	return err == nil, err
}

func GetUserID(c *gin.Context) (userId uint64, err error) {
	authHeader := c.Request.Header.Get("Authorization")
	claims, err := ParseToken(authHeader)
	if err != nil {
		return 0, err
	}

	userId = claims.UserID
	return
}

func GetUserinfo(c *gin.Context) (*entity.User, error) {
	uid, err := GetUserID(c)
	if err != nil {
		return nil, err
	}

	userInfo, err := user.GetUserByID(dao.Ctx(c), uid)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
