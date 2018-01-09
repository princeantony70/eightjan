package main()

import (
    _ "github.com/Go-SQL-Driver/MySQL"
    "database/sql"
    "fmt"
)

func main(){
  db,err :=sql.Open("mysql","roo:nfnt@tcp(127.0.0.1:3306)/gotest)
  if err != nil {
		fmt.Print(err.Error())
}

  stmt,err:=db.Prepare("INSERT person SET id=?,first_name=?,last_name=?")
  if err != nil {
		fmt.Print(err.Error())
}

  res, err := stmt.Exec(1, "prince", "antony")
  if err != nil {
		fmt.Print(err.Error())
}
}
