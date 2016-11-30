package sql
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
func main() {
	//sql.Open(). This returns an *sql.DB:
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:54469/D0018E-Database-Technology")
	if err != nil {
		fmt.Println(err)
		}
	defer db.Close()
}

func test(){

fmt.Println("heej")
}
