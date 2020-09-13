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

// DbInsert はデータを挿入する
func DbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず!")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

// DbUpdate はデータを更新する
func DbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず!")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// DbGetAll はデータをすべて取得する
func DbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず!")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// DbGetOne は指定したIDのデータを取り出す
func DbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず!")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

// DbDelete は指定したIDのデータを削除する
func DbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず!")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

// DbInit はデータベースを初期化する
func DbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず!")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}
