// @APIVersion 1.0.0
// @Title model
// @Description CRUD for books.
// @Contact tatink-t-s-r@yandex.ru
package model

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/goinggo/tracelog"
)

//** TYPES
type Book struct {
	Id     int64
	Name   string
	Author string
}

type Library struct {
	Id   int64
	Name string
}

type Book_Library struct {
	Id         int64
	Book_id    int64
	Library_id int64
}

func init() {
	orm.RegisterModel(new(Book), new(Library), new(Book_Library))
	orm.Debug = true
}

func GetBooksInLibrary(libraryId int64) (*[]Book, error) {
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

func AddBookInLibrary(libraryId int64, bookId int64) bool {
	o := orm.NewOrm()
	_, err := o.Raw(`
		INSERT INTO book_library(book_id, library_id) VALUES(?, ?)`, bookId, libraryId).Exec()
	if err != nil {
		tracelog.Errorf(err, "main", "model.AddBookInLibrary", "Failed to insert book")
		return false
	}
	return true
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
		return nil, err
	} else if err == orm.ErrMissPK {
		tracelog.Errorf(err, "main", "model.GetBook", "Failed to get book")
		return nil, err
	} else {
		return &book, err
	}
}

func AddBook(jsonStr []byte) bool {
	o := orm.NewOrm()
	var book Book
	err := json.Unmarshal(jsonStr, &book)
	if err != nil {
		tracelog.Errorf(err, "main", "model.AddBook", "Failed to insert json")
		return false
	}
	_, err = o.Insert(&book)
	if err != nil {
		tracelog.Errorf(err, "main", "model.AddBook", "Failed to insert book")
		return false
	}
	return true
}

func UpdateBook(jsonStr []byte) bool {
	o := orm.NewOrm()
	var book Book
	err := json.Unmarshal(jsonStr, &book)
	if err != nil {
		tracelog.Errorf(err, "main", "model.UpdateBook", "Failed to parse json")
		return false
	}
	_, err = o.Update(&book)
	if err != nil {
		tracelog.Errorf(err, "main", "model.UpdateBook", "Failed to update book")
		return false
	}
	return true
}

func DeleteBook(bookId int64) bool {
	o := orm.NewOrm()
	if _, err := o.Delete(&Book{Id: bookId}); err == nil {
		tracelog.Errorf(err, "main", "model.DeleteBook", "Failed to delete book")
		return false
	}
	return true
}

//Libraries
func GetLibraries() (*[]Library, error) {
	var rows []Library
	o := orm.NewOrm()
	_, err := o.Raw(` 
		SELECT id, name 
		FROM library`).QueryRows(&rows)
	if err != nil {
		tracelog.Errorf(err, "main", "model.GetBooks", "Failed to getbooks")
	}
	return &rows, err
}

func GetLibrary(libId int64) (*Library, error) {
	o := orm.NewOrm()
	lib := Library{Id: libId}
	err := o.Read(&lib)

	if err == orm.ErrNoRows {
		return nil, err
	} else if err == orm.ErrMissPK {
		tracelog.Errorf(err, "main", "model.GetBook", "Failed to get book")
		return nil, err
	} else {
		return &lib, err
	}
}

func AddLibrary(jsonStr []byte) bool {
	o := orm.NewOrm()
	var lib Library
	err := json.Unmarshal(jsonStr, &lib)
	if err != nil {
		tracelog.Errorf(err, "main", "model.AddBook", "Failed to insert json")
		return false
	}
	_, err = o.Insert(&lib)
	if err != nil {
		tracelog.Errorf(err, "main", "model.AddBook", "Failed to insert book")
		return false
	}
	return true
}

func UpdateLibrary(jsonStr []byte) bool {
	o := orm.NewOrm()
	var lib Library
	err := json.Unmarshal(jsonStr, &lib)
	if err != nil {
		tracelog.Errorf(err, "main", "model.UpdateBook", "Failed to parse json")
		return false
	}
	_, err = o.Update(&lib)
	if err != nil {
		tracelog.Errorf(err, "main", "model.UpdateBook", "Failed to update book")
		return false
	}
	return true
}

func DeleteLibrary(libId int64) bool {
	o := orm.NewOrm()
	if _, err := o.Delete(&Library{Id: libId}); err == nil {
		tracelog.Errorf(err, "main", "model.DeleteBook", "Failed to delete book")
		return false
	}
	return true
}
