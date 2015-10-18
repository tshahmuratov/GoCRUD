package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"books/models"
	"github.com/goinggo/tracelog"
	"strconv"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	libraries, num, _ := model.GetLibraries()
	if (num != 0) {
		books, err := model.GetBooks((*libraries)[0].Id)
	
		if (err == orm.ErrNoRows) {
			tracelog.Errorf(err, "main", "MainController", "No result found.")
		}
		c.Data["Libraries"] = libraries
		c.Data["Books"] = books
	}
	c.TplNames = "books.tpl"
	tracelog.Info("main", "main", "finish")
}

type BooksInLibraryContorller struct {
	beego.Controller
}

func (c *BooksInLibraryContorller) listBooksInLibrary() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":libId"), 10, 64)
	books, _ := model.GetBooks(id)
	c.Data["json"] = &books
    c.ServeJson()
}

func (c *BooksInLibraryContorller) addBookInLibrary() {
	//var datapoint Datapoint
	//book = json.Unmarshal(c.Ctx.Input.RequestBody, &datapoint)
	//res, err := model.AddBook(c.Ctx.Input.RequestBody)
	//prepareJson(c *BooksInLibraryContorller, res, err)
}

type BooksController struct {
	beego.Controller
}

func (c *BooksController) URLMapping() {
    c.Mapping("showAllBooks", c.showAllBooks)
    c.Mapping("showBook", c.showBook)
    c.Mapping("addBook", c.addBook)
    c.Mapping("updateBook", c.updateBook)
    c.Mapping("deleteBook", c.deleteBook)
}

func (c *BooksController) showAllBooks() {
	res, _ := model.GetAllBooks()
	if (res == nil) {
		c.Abort("404")
	}
	c.Data["json"] = &res
    c.ServeJson()
}

func (c *BooksController) showBook() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":bookId"), 10, 64)
	res, _ := model.GetBook(id)
	c.Data["json"] = &res
    c.ServeJson()
}

func (c *BooksController) addBook() {
	respByte, err := ioutil.ReadAll(c.Ctx.Request.Body)
    if err != nil {
		tracelog.Errorf(err, "main", "addBook", "Failed to read response data.")
		c.Abort("500")
		return
	}
	res := model.AddBook(respByte)
	if (!res) {
		c.Abort("500")
	}
	c.Data["json"] = &res
    c.ServeJson()
}

func (c *BooksController) updateBook() {
	respByte, err := ioutil.ReadAll(c.Ctx.Request.Body)
    if err != nil {
		tracelog.Errorf(err, "main", "updateBook", "Failed to read response data.")
		c.Abort("500")
		return
	}
	res := model.UpdateBook(respByte)
	if (!res) {
		c.Abort("500")
	}
	c.Data["json"] = &res
    c.ServeJson()
}

func (c *BooksController) deleteBook() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":bookId"), 10, 64)
	res := model.DeleteBook(id)
	if (!res) {
		c.Abort("500")
	}
	c.Data["json"] = &res
    c.ServeJson()
}

type LibraryController struct {
	beego.Controller
}

func (c *LibraryController) addLibrary() {
	
}

func (c *LibraryController) showAllLibraries() {
	
}

func (c *LibraryController) showLibrary() {
	
}

func (c *LibraryController) updateLibrary() {
	
}

func (c *LibraryController) deleteLibrary() {
	
}
//func (c *BooksController) PrepareJson