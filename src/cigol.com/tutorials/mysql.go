package	main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type User struct {
	id																	int
	username, password string
}
var sqldata map[interface{}]interface{}

func	main()	{
	var u User
	db, err := sql.Open("mysql", "root:smartpaas@tcp(10.126.3.163:3306)/paas_xuxy?charset=utf8")
	check(err)

	stmt, err := db.Prepare("INSERT	user SET username=?,password=?")
	check(err)

	res, err := stmt.Exec("xiaowei", "xiaowei")
	check(err)

	id, err := res.LastInsertId()
	check(err)
	fmt.Println(id)

	rows, err := db.Query("SELECT * FROM user")
	check(err)
	fmt.Println(rows.Columns())

	userinfo := make(map[interface{}]interface{})
	for rows.Next() {
		err := rows.Scan(&u.id,	&u.username, &u.password)
		check(err)
		userinfo[u.id] = u
	}
	fmt.Println(userinfo)
}
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}