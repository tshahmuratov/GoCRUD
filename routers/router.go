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
			beego.NSNamespace("/booksInLibrary",
				beego.NSNamespace("/:libId([0-9]+)",
					beego.NSRouter("/entries", &controllers.BooksInLibraryContorller{},"get:listBooksInLibrary"),
					beego.NSRouter("/entries", &controllers.BooksInLibraryContorller{},"post:addBookInLibrary"),
				),
			),
			beego.NSRouter("/books", &controllers.BooksController{},"get:showAllBooks"),
			beego.NSRouter("/books", &controllers.BooksController{},"put:updateBook"),
			beego.NSRouter("/books/:bookId([0-9]+)", &controllers.BooksController{},"get:showBook"),
			beego.NSRouter("/books/:bookId([0-9]+)", &controllers.BooksController{},"post:addBook"),
			beego.NSRouter("/books/:bookId([0-9]+)", &controllers.BooksController{},"delete:deleteBook"),
			/*beego.NSNamespace("/books",
			),*/
			beego.NSRouter("/libraries", &controllers.LibraryController{},"get:showAllLibraries"),
			beego.NSRouter("/libraries", &controllers.LibraryController{},"post:addLibrary"),
			beego.NSNamespace("/libraries",
				beego.NSRouter("/:libId([0-9]+)", &controllers.LibraryController{},"get:showLibrary"),
				beego.NSRouter("/:libId([0-9]+)", &controllers.LibraryController{},"put:updateLibrary"),
				beego.NSRouter("/:libId([0-9]+)", &controllers.LibraryController{},"delete:deleteLibrary"),
			),
		),
	)
		
	//register namespace
	beego.AddNamespace(ns)
}
