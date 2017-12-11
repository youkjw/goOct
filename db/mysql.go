package mysql

import(
	"database/sql"
	"fmt"
	"strconv"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func Insert(){
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go?charset=utf8")
	checkErr(err)
	defer db.Close()

	var count int
	err = db.QueryRow("select count(*) from user").Scan(&count)

	t := time.Now()
	count += 1;	
	var result sql.Result
	result, err = db.Exec("insert into user(name, create_time) values(?,?)", ("test" + strconv.Itoa(count)), t.Unix())
	checkErr(err)

	lastId, _ := result.LastInsertId() 
	fmt.Println("新插入记录的ID为", lastId)  
}

func Select(){
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go?charset=utf8")
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("select * from user")
	checkErr(err)

	var rows *sql.Rows  
	rows, err = stmt.Query()
	checkErr(err)

	for rows.Next(){
		var id int
		var name string
		var time int
		rows.Scan(&id, &name, &time)

		fmt.Println(id, name, time)
	}

	rows.Close()
}

func checkErr(err error){
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
