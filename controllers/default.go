package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"books/models"
	"github.com/goinggo/tracelog"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	tracelog.Info("main", "main", "GetBooks")
	libraries, num, _ := model.GetLibraries()
	if (num != 0) {
		books, _, err := model.GetBooks((*libraries)[0].Id)
	
		tracelog.Info("main", "main", "ok", books)
		if err == orm.ErrNoRows {
			fmt.Println("No result found.")
		} else if err == orm.ErrMissPK {
			fmt.Println("No primary key found.")
		}
		c.Data["Libraries"] = libraries
		c.Data["Books"] = books
	}
	c.TplNames = "books.tpl"
	tracelog.Info("main", "main", "finish")
}