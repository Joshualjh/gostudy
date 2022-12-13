package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = ""
	database = ""
	user     = ""
	password = ""
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	//db()
	r := setupRouter()
	r.Run(":8080")
}

func db() {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("successfully created connection to database.")

	sqlStatement, err := db.Prepare("INSERT INTO inventory (name, quantity) VALUES (?, ?);")
	res, err := sqlStatement.Exec("ddang", 150)
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	res, err = sqlStatement.Exec("doodnm", 154)
	checkError(err)
	rowCount, err = res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	res, err = sqlStatement.Exec("del", 100)
	checkError(err)
	rowCount, err = res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")

	var (
		id       int
		name     string
		quantity int
	)

	// Read some data from the table.
	rows, err := db.Query("SELECT id, name, quantity from inventory;")
	checkError(err)
	defer rows.Close()
	fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&id, &name, &quantity)
		checkError(err)
		fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
	}
	err = rows.Err()
	checkError(err)
	fmt.Println("Done.")

}
