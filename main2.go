package main

import "fmt"
import "database/sql"
import "os"
import _ "github.com/go-sql-driver/mysql"

var username, password, table, table_name, exit string

func connect() *sql.DB{
	var mariadb, err = sql.Open("mysql", username+":"+password+"@/"+table)
	err =mariadb.Ping
	if err != nil {
		fmt.Println("Failed to open database")
		fmt.Println("Database not connected")
		os.Exit(0)
			}
	return mariadb
			}

func outputsql() {
	var mariadb = connect()
	defer mariadb.Close()

	var Id_Mhs, Nama_Mhs string

	fmt.Println(" ")
	fmt.Println("Showing database :")
	fmt.Println(" ")

	rows, _ := db.Query("select * from " + table_name)

	for rows.Next() {
		rows.Scan(&Id_Mhs, &nama)
		fmt.Println("NIM  :" + Id_Mhs)
		fmt.Println("Nama :" + Nama_Mhs)
		fmt.Println(" ")

	}
	fmt.Println("Press enter to exit")
	fmt.Scanln(&exit)
}

func main() {
	fmt.Print("Masukkan user mysql     = ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password mysql = ")
	fmt.Scanln(&password)
	fmt.Print("Masukan nama database   = ")
	fmt.Scanln(&table)
	fmt.Print("Masukan nama tabel 	   = ")
	fmt.Scanln(&table_name)
	fmt.Println(" ")
	outputsql()

}