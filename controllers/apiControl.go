package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"movienotes/gosession"
	"movienotes/models"
	"movienotes/settings"
)

type KeyApiStr KeyApiS

type KeyApiS struct {
	Notes    []*models.Note
	Version  string
	KeyViews string
}

var page [5]string

func (c *MainController) ApiGetPage() {
	index, err := ioutil.ReadFile("views/api/index.tpl")
	if err != nil {
		fmt.Println(err)
	}
	addNote, err := ioutil.ReadFile("views/api/addNote.tpl")
	if err != nil {
		fmt.Println(err)
	}
	editNote, err := ioutil.ReadFile("views/api/editNote.tpl")
	if err != nil {
		fmt.Println(err)
	}
	keyNo, err := ioutil.ReadFile("views/api/keyNo.tpl")
	if err != nil {
		fmt.Println(err)
	}
	keyNull, err := ioutil.ReadFile("views/api/keyNull.tpl")
	if err != nil {
		fmt.Println(err)
	}
	page[0] = string(index)
	page[1] = string(addNote)
	page[2] = string(editNote)
	page[3] = string(keyNo)
	page[4] = string(keyNull)
	jsonPage, err := json.Marshal(page)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Ctx.ResponseWriter.Write(jsonPage)
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
}

func (c *MainController) ApiGetKey() {
	keyid := c.Ctx.Request.URL.Query().Get("keyid")
	version := c.Ctx.Request.URL.Query().Get("version")
	if keyid == "" {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	sess, err := gosession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("GETSESSION ERROR: ", err)
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	keyVersion := fmt.Sprint(settings.GetKeyVersion(keyid))
	if version == "" {
		switch keyVersion {
		case "-1":
			c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
			c.Ctx.ResponseWriter.Write([]byte("-1"))
		default:
			note := settings.GetNotes(keyid)
			key, err := ioutil.ReadFile("views/api/key.tpl")
			if err != nil {
				fmt.Println(err)
			}
			sourceTemplate, err := template.New("ET").Parse(string(key))
			if err != nil {
				panic(err)
			}
			var keyApiStr KeyApiStr
			keyApiStr.Notes = note
			keyApiStr.Version = keyVersion
			keyApiStr.KeyViews = keyid
			c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
			sourceTemplate.Execute(c.Ctx.ResponseWriter, keyApiStr)
		}
	} else {
		if keyVersion == version {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
			c.Ctx.ResponseWriter.Write([]byte(""))
		} else {
			switch keyVersion {
			case "-1":
				c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
				c.Ctx.ResponseWriter.Write([]byte("-1"))
			case "0":
				c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
				c.Ctx.ResponseWriter.Write([]byte("0"))
			default:
				note := settings.GetNotes(keyid)
				key, err := ioutil.ReadFile("views/api/key.tpl")
				if err != nil {
					fmt.Println(err)
				}
				sourceTemplate, err := template.New("ET").Parse(string(key))
				if err != nil {
					panic(err)
				}
				var keyApiStr KeyApiStr
				keyApiStr.Notes = note
				keyApiStr.Version = keyVersion
				keyApiStr.KeyViews = keyid
				c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
				sourceTemplate.Execute(c.Ctx.ResponseWriter, keyApiStr)
			}
		}
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	sess.Set("imgFile", "")
}
