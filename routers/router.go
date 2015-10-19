// @APIVersion 1.0.0
// @Title router
// @Description CRUD for books.
// @Contact tatink-t-s-r@yandex.ru
package routers

import (
	"books/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns :=
	beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			//books
			beego.NSRouter("/books", &controllers.BooksController{},"get:showAllBooks"),
			beego.NSRouter("/books", &controllers.BooksController{},"post:addBook"),
			beego.NSRouter("/books/:bookId([0-9]+)", &controllers.BooksController{},"get:showBook"),
			beego.NSRouter("/books/:bookId([0-9]+)", &controllers.BooksController{},"put:updateBook"),
			beego.NSRouter("/books/:bookId([0-9]+)", &controllers.BooksController{},"delete:deleteBook"),
			//libraries
			beego.NSRouter("/libraries", &controllers.LibraryController{},"get:showAllLibraries"),
			beego.NSRouter("/libraries", &controllers.LibraryController{},"post:addLibrary"),
			beego.NSRouter("/libraries/:libId([0-9]+)", &controllers.LibraryController{},"get:listBooksInLibrary"),
			beego.NSRouter("/libraries/:libId([0-9]+)/:bookId([0-9]+)", &controllers.LibraryController{},"post:addBookInLibrary"),
			beego.NSRouter("/libraries/:libId([0-9]+)", &controllers.LibraryController{},"put:updateLibrary"),
			beego.NSRouter("/libraries/:libId([0-9]+)", &controllers.LibraryController{},"delete:deleteLibrary"),
		),
	)
		
	//register namespace
	beego.AddNamespace(ns)
}
