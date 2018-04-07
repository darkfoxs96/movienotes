package models

import (
	"mime/multipart"
)

type Key struct {
	Key      string `form:"id,text,id:" valid:"MinSize(6);MaxSize(30)"`
	Password string `form:"password,text,password:" valid:"MinSize(6);MaxSize(30)"`
}

type Note struct {
	Id     int    `form:"idnote,number,idnote:"`
	Name   string `form:"name,text,name:" valid:"MaxSize(20)"`
	Note   string `form:"note,text,note:" valid:"MaxSize(255)"`
	Rating int    `form:"rating,number,rating:" valid:"Min(1);Max(5)"`
	KeyId  string
}

type AddNote struct {
	Img    multipart.File `form:"img,file,img:"`
	Name   string         `form:"name,text,name:" valid:"MaxSize(20)"`
	Note   string         `form:"note,text,note:" valid:"MaxSize(255)"`
	Rating int            `form:"rating,number,rating:" valid:"Min(1);Max(5)"`
	KeyId  string         `form:"idkey,text,idkey:"valid:"MaxSize(30)"`
}
