// @APIVersion 1.0.0
// @Title model
// @Description CRUD for books.
// @Contact tatink-t-s-r@yandex.ru
package model

import (
	"github.com/astaxie/beego/orm"
	"github.com/goinggo/tracelog"
)

//** TYPES

type Book struct {
	Id		int64
	Name 	string
	Author	string
}

type Library struct {
	Id		int64
	Name 	string
}

type Book_Library struct {
	Id				int64
	Book_id			*Book 		`orm:"rel(one)"`
	Library_id		*Library 	`orm:"rel(one)"`
}

func init() {
    orm.RegisterModel(new(Book), new(Library), new(Book_Library))
}

func GetLibraries() (*[]Library, int64, error) {
    orm.Debug = true
	var rows []Library
	o := orm.NewOrm()
	num, err := o.Raw(` 
		SELECT id, name 
		FROM library`).QueryRows(&rows)
	if err != nil {
		tracelog.Errorf(err, "main", "model.GetBooks", "Failed to getbooks")
	}
    return &rows, num, err
}

func GetBooks(libraryId int64) (*[]Book, int64, error) {
    orm.Debug = true
	var rows []Book
	o := orm.NewOrm()
	num, err := o.Raw(`
		SELECT b.id, b.name, b.author 
		FROM book b 
		join book_library con on b.id = con.book_id 
		join library s on con.library_id = s.id
		WHERE s.id = ?`, libraryId).QueryRows(&rows)
	if err != nil {
		tracelog.Errorf(err, "main", "model.GetBooks", "Failed to getbooks")
	}
    return &rows, num, err
}
