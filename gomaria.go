skysql:dellis$ cat gomaria.go 
package main 
import ( 
	"fmt" 
	"database/sql" 
	_ "github.com/go-sql-driver/mysql" 
) 
func main() { 
	// Create the database handle, confirm driver is present 
	db, _ := sql.Open("mysql", "dellis:@/shud") 
	defer db.Close() 

	// Connect and check the server version 
	var version string 
	db.QueryRow("SELECT VERSION()").Scan(&version) 
	fmt.Println("Connected to:", version) 
}

skysql:dellis$ go build gomaria.go 
skysql:dellis$ ./gomaria 
Connected to: 10.0.11-MariaDB
