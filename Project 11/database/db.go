package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)


var db *gorm.DB


type Bookmark struct {
	gorm.Model
	Name string `json:"name"`
	URL  string `json:"url"`
}

func InitDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.AutoMigrate(&Bookmark{})

	return nil
}

func CreateBookmark(name string, url string) (Bookmark, error) {
	bookmark := Bookmark{Name: name, URL: url}

	db.Create(&bookmark)

	return bookmark, nil
}

func GetAllBookmarks() ([]Bookmark, error) {
	var bookmarks []Bookmark

	db.Find(&bookmarks)

	return bookmarks, nil
}

func GetBookmark(id string) (Bookmark, error) {
	var bookmark Bookmark

	db.First(&bookmark, id)

	return bookmark, nil
}

func UpdateBookmarkPut(id string, name string, url string) (Bookmark, error) {
	var bookmark Bookmark

	if result := db.First(&bookmark, id); result.Error != nil {
		return bookmark, result.Error
		
	}

	bookmark.Name = name
	bookmark.URL = url
	db.Save(&bookmark)

	return bookmark, nil
}


func UpdateBookmarkPatch(id string, name string, url string) (Bookmark, error) {
	var bookmark Bookmark

	if result := db.First(&bookmark, id); result.Error != nil {
		return bookmark, result.Error
		
	}

	if result := db.Model(&bookmark).Updates(Bookmark{Name:name,URL:url}); result.Error != nil {
		return bookmark, result.Error
	}

	return bookmark, nil
}


func DeleteBookmark(id string) (int64, error) {
	db = db.Delete(&Bookmark{}, id)

	return db.RowsAffected, nil
}