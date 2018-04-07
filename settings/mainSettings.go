package settings

import (
	"fmt"

	"database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var (
	DataBase *sql.DB
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=135642 dbname=movie_notes sslmode=disable")
	orm.RunCommand()
	db, err := orm.GetDB("default")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	DataBase = db
}
