package dbaccessor

import (
	"github.com/jinzhu/gorm"

	// gorm を使用するために必要
	_ "github.com/mattn/go-sqlite3"
)

// Todo はタスク用の構造体
type Todo struct {
	gorm.Model
	Text   string
	Status string
}

func DbOpen() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず!")
	}
	db.LogMode(true)
	return db
}

// DbInsert はデータを挿入する
func DbInsert(db *gorm.DB, text string, status string) {
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

// DbUpdate はデータを更新する
func DbUpdate(db *gorm.DB, id int, text string, status string) {
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// DbGetAll はデータをすべて取得する
func DbGetAll(db *gorm.DB) []Todo {
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// DbGetOne は指定したIDのデータを取り出す
func DbGetOne(db *gorm.DB, id int) Todo {
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

// DbDelete は指定したIDのデータを削除する
func DbDelete(db *gorm.DB, id int) {
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

// DbInit はデータベースを初期化する
func DbInit(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
	defer db.Close()
}
