skysql:dellis$ cat gomaria.go 
package main 
import ( 
"sync" 
"database/sql" 
_ "github.com/go-sql-driver/mysql" 
) func main() {
 // Create the database handle, confirm driver is present 
db, _ := sql.Open("mysql", "dellis:@/shud") 
defer db.Close()

 // Test several connections 
var wg sync.WaitGroup 
for i := 0; i 
