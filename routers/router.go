package routers

import (
	"github.com/astaxie/beego"

	"movienotes/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/key/", &controllers.MainController{}, "get:GetKey")
	beego.Router("/key/createkey", &controllers.MainController{}, "post:AddKey")
	beego.Router("/key/createnote", &controllers.MainController{}, "post:AddNote")
	beego.Router("/key/createnote/loadimgapi", &controllers.MainController{}, "get:LoadImgApi")
	beego.Router("/key/editnote", &controllers.MainController{}, "put:EditNote")
	beego.Router("/key/deletenote", &controllers.MainController{}, "delete:DeleteNote")

	beego.Router("/user/outkey", &controllers.MainController{}, "get:KeyOut")
	beego.Router("/user/inkey", &controllers.MainController{}, "post:KeyIn")

	beego.Router("/api/page", &controllers.MainController{}, "get:ApiGetPage")
	beego.Router("/api/key/", &controllers.MainController{}, "get:ApiGetKey")
}
