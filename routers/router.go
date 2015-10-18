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
}
