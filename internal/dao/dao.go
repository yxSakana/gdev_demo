package dao

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/yxSakana/gdev_demo/settings"
)

var mysqlDB *gorm.DB

func init() {
	cfg := settings.Settings.Database.Mysql
	link := cfg.Link
	if link == "" {
		link = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.User, cfg.Password, cfg.Address, cfg.Dbname)
	}

	var err error
	mysqlDB, err = gorm.Open(mysql.Open(link), &gorm.Config{})
	if err != nil {
		log.Printf("mysql connect err: %v", err)
		panic(err)
	}
}

func DBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", mysqlDB)
		c.Next()
	}
}

func Ctx(c *gin.Context) *gorm.DB {
	db, exists := c.Get("DB")
	if !exists {
		panic("DB not found in Context")
	}
	return db.(*gorm.DB)
}
