package main


import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:nfn@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}
	stmt, err := db.Prepare("ALTER TABLE person ADD(emailid  varchar(30)")
	if err != nil {
		fmt.Println(err.Error())
}
_, err = stmt.Exec()
	if err != nil {
fmt.Println(err.Error())
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Person struct {
		Id         int
		First_Name string
		Last_Name  string
	}
	router := gin.Default()
	// Add API handlers her
  // POST new person details
  	router.POST("/person", func(c *gin.Context) {
  		var buffer bytes.Buffer
  		first_name := c.PostForm("first_name")
  		last_name := c.PostForm("last_name")
  		stmt, err := db.Prepare("INSERT INTO person (first_name, last_name) values(?,?);")
  		if err != nil {
  			fmt.Print(err.Error())
  		}
  		_, err = stmt.Exec(first_name, last_name)

  		if err != nil {
  			fmt.Print(err.Error())
  		}

  		// Fastest way to append strings
  		buffer.WriteString(first_name)
  		buffer.WriteString(" ")
  		buffer.WriteString(last_name)
  		defer stmt.Close()
  		name := buffer.String()
  		c.JSON(http.StatusOK, gin.H{
  			"message": fmt.Sprintf(" %s successfully created", name),
  		})
  })
	router.Run(":7070")
}}
