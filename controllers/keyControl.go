package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"

	"movienotes/gosession"
	"movienotes/models"
	"movienotes/settings"
)

func (c *MainController) GetKey() {
	sess, err := gosession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("GETSESSION ERROR: ", err)
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	c.navbar(sess)
	keyid := c.Ctx.Request.URL.Query().Get("keyid")
	if keyid == "" {
		c.TplName = "body/keyNo.tpl"
		return
	}
	note := settings.GetNotes(keyid)
	if !settings.IsKey(keyid) {
		c.Data["CreateKey"] = keyid
		c.TplName = "body/keyNo.tpl"
		return
	}
	c.Data["Notes"] = note
	c.Data["KeyViews"] = keyid
	c.TplName = "body/key.tpl"
	sess.Set("imgFile", "")
}

func (c *MainController) AddKey() {
	sess, err := gosession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("GETSESSION ERROR: ", err)
	}
	var key models.Key
	c.ParseForm(&key)
	valid := validation.Validation{}
	isValid, _ := valid.Valid(&key)
	if isValid {
		err = settings.AddKey(key.Key, key.Password)
		if err != nil {
			if err.Error() == "the key is busy" {
				c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
				return
			}
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			return
		}
		sess.Set("IsAutorized", true)
		sess.Set("KeyId", key.Key)
		c.CruSession = sess
		sess.SessionRelease(c.Ctx.ResponseWriter)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	} else {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	}
}
