package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	search_data "github.com/gagaspbahar/dna-pattern-matching-web/algo/search"
	"github.com/gin-gonic/gin"
)

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "tubes3"
)

type Penyakit struct {
	Penyakit string `json:"penyakit"`
	Content  string `json:"content"`
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func main() {
	r := gin.Default()
	db, err := sql.Open("mysql", dsn())

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()
	search_data.Search_db(db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.POST("/add", postAddPenyakit)
	r.POST("/upload/:filename", handleUpload)
	r.Run()
}

func handleUpload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	header_filename := header.Filename
	log.Printf("filename: %s", header_filename)
	filename := c.Param("filename")
	filename = filename + ".txt"
	fmt.Println(filename)
	out, err := os.Create("/data/sequence" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "/data/sequence/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}
