package main

import (
	"fmt"
	"gpapproval/db/config"
)

// 文章
type Topic struct {
	Id         int        `gorm:"primary_key"`
	Title      string     `gorm:"not null"`
	UserId     int        `gorm:"not null"`
	CategoryId int        `gorm:"not null"`
	//Category   Category `gorm:"foreignkey:CategoryId"`//文章所属分类外键
	//User       Users      `gorm:"foreignkey:UserId"`//文章所属用户外键
}

// 用户
type User struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

type Content struct {
	Info string `gorm:"info"`
	Author string `gorm:"author"`
}
// 分类
type Category struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"not null"`
	Content
}

func main() {
	config.GetApprovalDbConfig("test")
	config.GetDb()
	//config.ClearAllData()
	//
	//if err := config.GetDb().CreateTable(Topic{}).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//if err := config.GetDb().CreateTable(User{}).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//for i:=0; i<4; i++ {
	//	title := fmt.Sprintf("ha%d", i)
	//	config.GetDb().Model(&Topic{}).Create(&Topic{Title:title, UserId:i})
	//	name := fmt.Sprintf("ji%d", i)
	//	config.GetDb().Model(&User{}).Create(&User{Name:name})
	//}
	ret := []Topic{}
	config.GetDb().Select("title").Joins("inner join users on users.id=topics.user_id where users.id>?",1).Offset(0).Limit(1).Find(&ret)
	fmt.Println(ret)

	//config.GetDb().Model(&Topic{}).Delete(Topics{Id:5})

	//if err := config.GetDb().CreateTable(Category{}).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}
}