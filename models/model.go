// @APIVersion 1.0.0
// @Title model
// @Description CRUD for books.
// @Contact tatink-t-s-r@yandex.ru
package model

import (
	"github.com/astaxie/beego/orm"
	"github.com/goinggo/tracelog"
	"encoding/json"
	"fmt"
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
	orm.Debug = true
}

func GetLibraries() (*[]Library, int64, error) {
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

func GetBooks(libraryId int64) (*[]Book, error) {
    var rows []Book
	o := orm.NewOrm()
	_, err := o.Raw(`
		SELECT b.id, b.name, b.author 
		FROM book b 
		join book_library con on b.id = con.book_id 
		join library s on con.library_id = s.id
		WHERE s.id = ?`, libraryId).QueryRows(&rows)
	if err != nil {
		tracelog.Errorf(err, "main", "model.GetBooks", "Failed to getbooks")
	}
    return &rows, err
}

func GetAllBooks() (*[]Book, error) {
    var rows []Book
	o := orm.NewOrm()
	_, err := o.Raw(`
		SELECT b.id, b.name, b.author 
		FROM book b`).QueryRows(&rows)
	if err != nil {
		tracelog.Errorf(err, "main", "model.GetBooks", "Failed to getbooks")
	}
    return &rows, err
}


func GetBook(bookId int64) (*Book, error) {
    o := orm.NewOrm()
	book := Book{Id: bookId}
	err := o.Read(&book)
	
	if err == orm.ErrNoRows {
		return nil, err;
	} else if err == orm.ErrMissPK {
		tracelog.Errorf(err, "main", "model.GetBook", "Failed to get book")
		return nil, err;
	} else {
		return &book, err
	}
}

func AddBook(jsonStr []byte) (bool) {
	o := orm.NewOrm()
	var book Book
	err := json.Unmarshal(jsonStr, &book)
   	if (err != nil) {
		tracelog.Errorf(err, "main", "model.AddBook", "Failed to insert json")
		return false
	}
	num, err := o.Insert(&book);
	if (err != nil) {
		tracelog.Errorf(err, "main", "model.AddBook", "Failed to insert book")
		return false
	}
	if (num != 1) {
		tracelog.Errorf(err, "main", "model.AddBook", fmt.Sprint("Failed to insert book num = ",num))
		return false
	}
	return true;
}

func UpdateBook(jsonStr []byte) (bool) {
	o := orm.NewOrm()
	var book Book
	err := json.Unmarshal(jsonStr, &book)
   	if (err != nil) {
		tracelog.Errorf(err, "main", "model.UpdateBook", "Failed to parse json")
		return false
	}
	num, err := o.Update(&book);
	if (err != nil) {
		tracelog.Errorf(err, "main", "model.UpdateBook", "Failed to update book")
		return false
	}
	if (num != 1) {
		tracelog.Errorf(err, "main", "model.UpdateBook", fmt.Sprint("Failed to update book num = ",num))
		return false
	}
	return true;
}

func DeleteBook(bookId int64) (bool) {
	o := orm.NewOrm()
	if _, err := o.Delete(&Book{Id: 1}); err == nil {
		tracelog.Errorf(err, "main", "model.DeleteBook", "Failed to delete book")
		return false
	}
	return true
}