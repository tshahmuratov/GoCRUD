package controllers

import (
	"books/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/goinggo/tracelog"
	"io/ioutil"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplNames = "books.tpl"
	tracelog.Info("main", "main", "finish")
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
	if res == nil {
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
	if !res {
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
	if !res {
		c.Abort("500")
	}
	c.Data["json"] = &res
	c.ServeJson()
}

func (c *BooksController) deleteBook() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":bookId"), 10, 64)
	res := model.DeleteBook(id)
	if !res {
		c.Abort("500")
	}
	c.Data["json"] = &res
	c.ServeJson()
}

type LibraryController struct {
	beego.Controller
}

func (c *LibraryController) URLMapping() {
	c.Mapping("addLibrary", c.addLibrary)
	c.Mapping("showAllLibraries", c.showAllLibraries)
	c.Mapping("showLibrary", c.showLibrary)
	c.Mapping("updateLibrary", c.updateLibrary)
	c.Mapping("deleteLibrary", c.deleteLibrary)
	c.Mapping("listBooksInLibrary", c.listBooksInLibrary)
	c.Mapping("addBookInLibrary", c.addBookInLibrary)
}

func (c *LibraryController) addLibrary() {
	respByte, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		tracelog.Errorf(err, "main", "addBook", "Failed to read response data.")
		c.Abort("500")
		return
	}
	res := model.AddLibrary(respByte)
	if !res {
		c.Abort("500")
	}
	c.Data["json"] = &res
	c.ServeJson()
}

func (c *LibraryController) showAllLibraries() {
	res, _ := model.GetLibraries()
	if res == nil {
		c.Abort("404")
	}
	c.Data["json"] = &res
	c.ServeJson()
}

func (c *LibraryController) showLibrary() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":bookId"), 10, 64)
	res, _ := model.GetLibrary(id)
	c.Data["json"] = &res
	c.ServeJson()
}

func (c *LibraryController) updateLibrary() {
	respByte, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		tracelog.Errorf(err, "main", "updateBook", "Failed to read response data.")
		c.Abort("500")
		return
	}
	res := model.UpdateLibrary(respByte)
	if !res {
		c.Abort("500")
	}
	c.Data["json"] = &res
	c.ServeJson()
}

func (c *LibraryController) deleteLibrary() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":bookId"), 10, 64)
	res := model.DeleteLibrary(id)
	if !res {
		c.Abort("500")
	}
	c.Data["json"] = &res
	c.ServeJson()
}

func (c *LibraryController) listBooksInLibrary() {
	libId, _ := strconv.ParseInt(c.Ctx.Input.Param(":libId"), 10, 64)
	books, err := model.GetBooksInLibrary(libId)

	if err == orm.ErrNoRows {
		tracelog.Errorf(err, "main", "MainController", "No result found.")
	}
	c.Data["json"] = &books
	c.ServeJson()
}

func (c *LibraryController) addBookInLibrary() {
	libId, _ := strconv.ParseInt(c.Ctx.Input.Param(":libId"), 10, 64)
	bookId, _ := strconv.ParseInt(c.Ctx.Input.Param(":bookId"), 10, 64)
	res := model.AddBookInLibrary(libId, bookId)
	if !res {
		c.Abort("500")
	}
	c.Data["json"] = &res
	c.ServeJson()
}
