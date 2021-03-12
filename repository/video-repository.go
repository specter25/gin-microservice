package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/specter25/gin-microservice/entity"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("mysql", "specter25dbms:Jio@1234@/gin-microservice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("failed to close database")
	}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}
func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}
func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}
func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos

}
