package mysql

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Insert(){
	db, err := sql.Open("mysql", "caoweiping:caoweiping@tcp(localhost:3306)/gamedata_tian?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connet mysql success")

	defer db.Close()
}
