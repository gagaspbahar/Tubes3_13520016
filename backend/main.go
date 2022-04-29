package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gagaspbahar/dna-pattern-matching-web/algo/add/stringmatching"
	search_data "github.com/gagaspbahar/dna-pattern-matching-web/algo/search"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "dlUn61tF6K"
	password = "RcEHz7Svaw"
	hostname = "remotemysql.com:3306"
	dbname   = "dlUn61tF6K"
)

type Penyakit struct {
	Penyakit string `json:"penyakit"`
}

type TesDNA struct {
	Nama        string `json:"nama"`
	SequenceDNA string `json:"sequenceDNA"`
	Penyakit    string `json:"penyakit"`
	Metode      string `json:"metode"`
}

type HistoryDNA struct {
	Query string `json:"query"`
}

type Record struct {
	Tanggal       string  `json:"tanggal"`
	Nama_pengguna string  `json:"nama_pengguna"`
	Nama_penyakit string  `json:"nama_penyakit"`
	Similarity    float64 `json:"similarity"`
	Status_tes    int     `json:"status_tes"`
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	db, err := sql.Open("mysql", dsn())

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/add", postAddPenyakit(db))
	r.POST("/upload/:filename", handleUpload)
	r.POST("/uploadUser/:filename", handleUploadUserSequence)
	r.POST("/tes", getTesDNA(db))
	r.POST("/history", getDNAHistory(db))
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func handleUpload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	filename := header.Filename
	log.Printf("filename: %s", filename)

	out, err := os.Create("data/sequence/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "/data/sequence/" + filename
	content, err := os.ReadFile("data/sequence/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	result := CheckDNASequence(string(content))
	if result {
		c.JSON(http.StatusOK, gin.H{"filepath": filepath, "message": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "failed"})
	}
}

func handleUploadUserSequence(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	filename := header.Filename
	log.Printf("filename: %s", filename)

	out, err := os.Create("data/sequenceUser/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "/data/sequenceUser/" + filename
	content, err := os.ReadFile("data/sequenceUser/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	result := CheckDNASequence(string(content))
	if result {
		c.JSON(http.StatusOK, gin.H{"filepath": filepath, "message": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "failed"})
	}
}

func postAddPenyakit(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var penyakit Penyakit
		// rawdata, _ := c.GetRawData()
		// fmt.Println(string(rawdata))
		if err := c.BindJSON(&penyakit); err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			log.Fatal(err)
			return
		}
		filename := penyakit.Penyakit + ".txt"
		content, err := os.ReadFile("data/sequence/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		result := CheckDNASequence(string(content))
		if result {
			insert, err := db.Query("INSERT INTO sequence_penyakit VALUES ( ?, ? )", penyakit.Penyakit, content)
			if err != nil {
				log.Printf("Error %s when insert to DB\n", err)
			}
			defer insert.Close()
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "failed"})
		}
	}
}

func getTesDNA(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var tesDNA TesDNA
		var result int
		var sim float64
		// rawdata, _ := c.GetRawData()
		// log.Printf("tesDNA: %s", rawdata)
		if err := c.BindJSON(&tesDNA); err != nil {
			log.Printf("gagal bind")
			c.String(http.StatusBadRequest, "Bad request")
			log.Fatal(err)
			return
		}

		filepath := "data/sequenceUser/" + tesDNA.SequenceDNA + ".txt"

		sequence, err := os.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
		}
		sequenceString := string(sequence)

		filepath = "data/sequence/" + tesDNA.Penyakit + ".txt"
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

		if result == 1 {
			sim = 100
		} else {
			sim = float64(stringmatching.LCS(sequenceString, spString)) / float64(len(spString)) * 100
			if sim >= 80 {
				result = 1
			}
		}

		tanggal := time.Now().Format("2006-01-02")

		insert, err := db.Query("INSERT INTO data_uji (tanggal_tes, nama_pengguna, nama_penyakit, similarity, status_tes) VALUES ( ?, ? , ?, ?, ? )", tanggal, tesDNA.Nama, tesDNA.Penyakit, sim, result)
		if err != nil {
			log.Printf("Error %s when insert to DB\n", err)
		}

		defer insert.Close()

		var record Record
		record.Nama_pengguna = tesDNA.Nama
		record.Nama_penyakit = tesDNA.Penyakit
		record.Similarity = sim
		record.Status_tes = result
		record.Tanggal = tanggal

		c.JSON(http.StatusOK, record)
	}
}

func getDNAHistory(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var results *sql.Rows
		var err error
		var history HistoryDNA
		var records []Record

		if err := c.BindJSON(&history); err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			log.Fatal(err)
			return
		}

		data := history.Query
		date, penyakit := search_data.CheckSequence(data)
		date, penyakit = search_data.Clear_whitespace(date, penyakit)
		if date != "" && penyakit != "" {
			results, err = db.Query("SELECT tanggal_tes, nama_pengguna, nama_penyakit, similarity, status_tes FROM data_uji WHERE tanggal_tes = ? AND nama_penyakit = ?", date, penyakit)
		} else if penyakit != "" {
			results, err = db.Query("SELECT tanggal_tes, nama_pengguna, nama_penyakit, similarity, status_tes FROM data_uji WHERE nama_penyakit = ?", penyakit)
		} else if date != "" {
			results, err = db.Query("SELECT tanggal_tes, nama_pengguna, nama_penyakit, similarity, status_tes FROM data_uji WHERE tanggal_tes = ?", date)
		} else {
			return
		}
		if err != nil {
			log.Printf("Error %s when SELECT from DB\n", err)
		}
		for results.Next() {
			var record Record
			// for each row, scan the result into our tag composite object
			err = results.Scan(&record.Tanggal, &record.Nama_pengguna, &record.Nama_penyakit, &record.Similarity, &record.Status_tes)
			if err != nil {
				log.Printf("Error %s when insert to DB\n", err)
			}
			records = append(records, record)
		}
		recordsJson, _ := json.Marshal(records)
		fmt.Println(records)
		fmt.Println(string(recordsJson))
		c.JSON(http.StatusOK, gin.H{"records": string(recordsJson)})
	}
}

func CheckDNASequence(input string) bool {
	res1, _ := regexp.MatchString(`^[CAGT]+$`, input)
	return res1
}
