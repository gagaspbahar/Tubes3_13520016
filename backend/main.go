package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gagaspbahar/dna-pattern-matching-web/algo/add/stringmatching"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

type TesDNA struct {
	Nama        string `json:"nama"`
	SequenceDNA string `json:"sequenceDNA"`
	Penyakit    string `json:"penyakit"`
	Metode      string `json:"metode"`
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
	// search_data.Search_db(db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.POST("/add", postAddPenyakit)
	r.POST("/upload/:filename", handleUpload)
	r.GET("/tes", getTesDNA)
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

func getTesDNA(c *gin.Context) {
	var tesDNA TesDNA
	var result int
	if err := c.BindJSON(&tesDNA); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		log.Fatal(err)
		return
	}

	filepath := "/data/sequenceUser/" + tesDNA.SequenceDNA + ".txt"

	sequence, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	sequenceString := string(sequence)

	filepath = "/data/sequence/" + tesDNA.Penyakit + ".txt"
	sequence_penyakit, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	spString := string(sequence_penyakit)

	if tesDNA.Metode == "KMP" {
		result = stringmatching.Boyermoore(sequenceString, spString)
	} else {
		result = stringmatching.KMP(sequenceString, spString)
	}

	tanggal := time.Now().Format("2006-01-02")

	c.JSON(http.StatusOK, gin.H{"result": result, "nama": tesDNA.Nama, "penyakit": tesDNA.Penyakit, "tanggal": tanggal})
}
