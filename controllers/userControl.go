package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"

	"movienotes/gosession"
	"movienotes/models"
	"movienotes/settings"
)

func (c *MainController) KeyIn() {
	sess, err := gosession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("GETSESSION ERROR: ", err)
	}
	var key models.Key
	c.ParseForm(&key)
	valid := validation.Validation{}
	isValid, _ := valid.Valid(&key)
	if isValid {
		if settings.IsKeyPassword(key.Key, key.Password) {
			sess.Set("IsAutorized", true)
			sess.Set("KeyId", key.Key)
			c.CruSession = sess
			sess.SessionRelease(c.Ctx.ResponseWriter)
			c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
			return
		} else {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		}
	} else {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	}
}

func (c *MainController) KeyOut() {
	sess, err := gosession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("GETSESSION ERROR: ", err)
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	sess.Set("IsAutorized", false)
	sess.Set("KeyId", "")
	c.CruSession = sess
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
}
