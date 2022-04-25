package main

import (
	"database/sql"
	"fmt"
	"log"
    _ "github.com/gagaspbahar/dna-pattern-matching-web/algo/add"
    "github.com/gagaspbahar/dna-pattern-matching-web/algo/search"
	_ "github.com/go-sql-driver/mysql"
)

const (  
    username = "root"
    password = ""
    hostname = "127.0.0.1:3306"
    dbname   = "tubes3_basdat"
)

func dsn() string {  
    return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func main() {
    db, err := sql.Open("mysql", dsn())
    if err != nil {
        log.Printf("Error %s when opening DB\n", err)
        return
    }
    defer db.Close()
    search_data.Search_db(db)
}