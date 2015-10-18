package main

import (
	_ "books/routers"
	_ "github.com/lib/pq"
	"os"
	"bytes"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/goinggo/tracelog"
	"github.com/astaxie/beego/config"
)

type Config struct {
    User []string
    Pass []string
    DBName []string
}

func main() {
	tracelog.StartFile(tracelog.LevelTrace, "logs", 1)
	tracelog.Started("main", "Initializing Postgres")
	var confName string = os.Getenv("BOOKS_CONF");
	if len(confName) == 0 { confName = "default" }
	s := []string{"conf/db/", confName, ".json"};
    jsonconf, err := config.NewConfig("json", strings.Join(s, ""))
	if err != nil {
		tracelog.Errorf(err, "main", "config.NewConfig", "Failed to find config", strings.Join(s, ""))
	}
	var User string = jsonconf.String("User")
	var Pass string = jsonconf.String("Pass")
	var DBName string = jsonconf.String("DBName")
	var buffer bytes.Buffer
	buffer.WriteString("user=")
	buffer.WriteString(User)
	buffer.WriteString(" password=")
	buffer.WriteString(Pass)
	buffer.WriteString(" dbname=")
	buffer.WriteString(DBName)
	s2 := []string{"user=", User, " password=", Pass, " dbname=", DBName, " sslmode=disable"};
	orm.RegisterDriver("postgres", orm.DR_Postgres)
    orm.RegisterDataBase("default", "postgres", strings.Join(s2, ""))
    beego.SetStaticPath("/images","static/images")
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/js","static/js")
    beego.SetStaticPath("/fonts","static/fonts")
    beego.EnableAdmin = true
    orm.RunCommand()
	beego.Run()
}

