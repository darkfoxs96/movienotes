package controllers

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/nfnt/resize"

	"movienotes/gosession"
	"movienotes/models"
	"movienotes/settings"
)

func (c *MainController) DeleteNote() {
	sess, err := gosession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("GETSESSION ERROR: ", err)
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	if sess.Get("IsAutorized") == false {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	if sess.Get("IsAutorized") == true {
		noteid := c.Ctx.Request.URL.Query().Get("idnote")
		id, err := strconv.Atoi(noteid)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			return
		}
		err = settings.DeleteNote(fmt.Sprint(sess.Get("KeyId")), id)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			return
		}
		c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
		settings.EditKeyVersion(fmt.Sprint(sess.Get("KeyId")), "remove")
		err = os.Remove("./static/content/" + c.Ctx.Request.URL.Query().Get("idnote") + ".jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
}

func (c *MainController) EditNote() {
	sess, err := gosession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("GETSESSION ERROR: ", err)
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	if sess.Get("IsAutorized") == false {
		c.Ctx.ResponseWriter.Write([]byte("Авторизуйтесь!"))
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	if sess.Get("IsAutorized") == true {
		var note models.Note
		c.ParseForm(&note)
		valid := validation.Validation{}
		isValid, _ := valid.Valid(&note)
		if isValid {
			note.KeyId = fmt.Sprint(sess.Get("KeyId"))
			err = settings.EditNote(note)
			if err != nil {
				if err.Error() == "Not your key" {
					c.Ctx.ResponseWriter.Write([]byte("Вы авторизованны под другим ключём!"))
					c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
					return
				}
				c.Ctx.ResponseWriter.Write([]byte(err.Error()))
				c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
				return
			}
			c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
			settings.EditKeyVersion(fmt.Sprint(sess.Get("KeyId")), "edit,add")
			return
		} else {
			c.Ctx.ResponseWriter.Write([]byte("Введённые данные не корректны!"))
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (c *MainController) AddNote() {
	sess, err := gosession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("GETSESSION ERROR: ", err)
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	if sess.Get("IsAutorized") == false {
		c.Ctx.ResponseWriter.Write([]byte("Ошибка доступа!"))
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	if sess.Get("IsAutorized") == true {
		keyid := fmt.Sprint(sess.Get("KeyId"))
		var note models.AddNote
		c.ParseForm(&note)
		valid := validation.Validation{}
		isValid, _ := valid.Valid(&note)
		if isValid {
			if keyid != note.KeyId {
				c.Ctx.ResponseWriter.Write([]byte("Ошибка доступа!"))
				c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
				return
			}
			c.Ctx.Request.ParseMultipartForm(2097152)
			file, _, err := c.Ctx.Request.FormFile("img")
			if err != nil {
				fmt.Println(err)
			}
			err, id := settings.AddNote(note)
			if err != nil {
				c.Ctx.ResponseWriter.Write([]byte(err.Error()))
				c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
				return
			}
			if file != nil {
				img, _, err := image.Decode(file)
				if err != nil {
					fmt.Println(err)
					c.Ctx.ResponseWriter.Write([]byte(err.Error()))
					c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
				}
				img = resize.Resize(400, 260, img, resize.Bicubic) // <-- Собственно изменение размера картинки
				imgOut, err := os.Create(fmt.Sprintf("./static/content/%d%s", id, ".jpg"))
				if err != nil {
					fmt.Println(err)
					c.Ctx.ResponseWriter.Write([]byte(err.Error()))
					c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
				}
				jpeg.Encode(imgOut, img, nil)
				file.Close()
				imgOut.Close()
				err = os.Remove("./static/content/" + fmt.Sprint(sess.Get("imgFile")))
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				err = os.Rename("./static/content/"+fmt.Sprint(sess.Get("imgFile")), fmt.Sprintf("./static/content/%d%s", id, ".jpg"))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			sess.Set("imgFile", "")
			c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
			settings.EditKeyVersion(fmt.Sprint(sess.Get("KeyId")), "edit,add")
			return
		} else {
			c.Ctx.ResponseWriter.Write([]byte("Введённые данные не корректны!"))
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func (c *MainController) LoadImgApi() {
	sess, err := gosession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("GETSESSION ERROR: ", err)
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	name := c.Ctx.Request.URL.Query().Get("name")
	url := GetImgUrl(name)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	}
	img, _, err := image.Decode(response.Body)
	if err != nil {
		fmt.Println(err)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	}
	img = resize.Resize(400, 260, img, resize.Bicubic) // <-- Собственно изменение размера картинки
	id := fmt.Sprintf("time%d", time.Now().Unix()*10)
	imgOut, err := os.Create("./static/content/" + id + ".jpg")
	if err != nil {
		fmt.Println(err)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	}
	jpeg.Encode(imgOut, img, nil)
	response.Body.Close()
	imgOut.Close()
	if sess.Get("imgFile") != nil && sess.Get("imgFile") != "" {
		err = os.Remove("./static/content/" + fmt.Sprint(sess.Get("imgFile")))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	sess.Set("imgFile", id+".jpg")
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(id + ".jpg"))
}

func GetImgUrl(searchTerm string) string {
	const endpoint = "https://api.cognitive.microsoft.com/bing/v7.0/images/search"
	token := "6b3b2d72307f4f9b8457b3442177ee19"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println(err)
	}
	param := req.URL.Query()
	param.Add("q", searchTerm)
	req.URL.RawQuery = param.Encode()
	req.Header.Add("Ocp-Apim-Subscription-Key", token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	ans := string(body)
	st := strings.Index(ans, `"contentUrl"`) + 15
	ans = ans[st:]
	st = strings.Index(ans, `"`)
	ans = ans[:st]
	ans = strings.Replace(ans, `\`, "", -1)
	return ans
}
