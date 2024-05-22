package corm

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"testing"
)

type Book struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	//CategoryID uint      `json:"categoryId"`
	//Category   *Category `json:"category"`
}

func TestNewGormDao(t *testing.T) {
	dns := "root:123456@tcp(127.0.0.1:3306)/gorse_demo?charset=utf8&parseTime=True&loc=Local"
	dao := NewClient().WithConfig(&Config{
		Dsn: dns,
	}).Open()
	err := dao.Db.AutoMigrate(&Book{})
	hlog.Debug(err)
}
