package add 

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"github.com/gagaspbahar/dna-pattern-matching-web/algo/add/stringmatching"
	"os"
)

func AddDataUji(db *sql.DB) {
    var filename string
    var penyakit string
	var pengguna string
	fmt.Println("Masukan nama pengguna: ")
    fmt.Scanln(&pengguna)
    fmt.Println("Masukan sequence: ")
    fmt.Scanln(&filename)
    fmt.Println("Masukan nama penyakit: ")
    fmt.Scanln(&penyakit)
    content, err := os.ReadFile("sequence/" + filename)
    if err != nil {
        log.Fatal(err)
    }
    result := CheckSequence(string(content))
    if (result) {
		curDate := time.Now()
		var seq_check string
		err := db.QueryRow("SELECT sequence FROM sequence_penyakit WHERE penyakit = ?", penyakit).Scan(&seq_check)
		if err != nil {
            log.Printf("Error %s when SELECT from DB\n", err)
        }
		result_uji := boyermoore.Boyermoore(string(content), seq_check)
        insert, err := db.Query("INSERT INTO data_uji (nama_pengguna, nama_penyakit, status_tes, tanggal_tes) VALUES ( ?, ? , ?, ? )", pengguna, penyakit, result_uji, curDate.Format("2006-01-02"))
        if err != nil {
            log.Printf("Error %s when insert to DB\n", err)
        }
        defer insert.Close()
    }
}