package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"

	"movienotes/gosession"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	sess, err := gosession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("GETSESSION ERROR: ", err)
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	c.navbar(sess)
	c.TplName = "body/index.tpl"
	sess.Set("imgFile", "")
}

func (c *MainController) navbar(sess session.Store) {
	if sess.Get("IsAutorized") == true {
		c.Data["Key"] = sess.Get("KeyId")
		c.Data["FunKey"] = "out"
		c.Data["IconKey"] = `sign-out`
	} else {
		c.Data["Key"] = "Нет"
		c.Data["FunKey"] = "in"
		c.Data["IconKey"] = `key`
	}
}
