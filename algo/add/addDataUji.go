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
    var filename, penyakit, pengguna, seq_check string
    var algo, result_uji int
    var sim float64
	fmt.Print("Masukan nama pengguna: ")
    fmt.Scanln(&pengguna)
    fmt.Print("Masukan nama file: ")
    fmt.Scanln(&filename)
    fmt.Print("Masukan nama penyakit: ")
    fmt.Scanln(&penyakit)
    content, err := os.ReadFile("data/sequenceUser/" + filename)
    if err != nil {
        log.Fatal(err)
    }
    result := CheckSequence(string(content))
    if (result) {
		err := db.QueryRow("SELECT sequence FROM sequence_penyakit WHERE penyakit = ?", penyakit).Scan(&seq_check)
		if err != nil {
            log.Printf("Error %s when SELECT from DB\n", err)
        }
        fmt.Println("1. Boyer-moore")
        fmt.Println("2. KMP")
        fmt.Print("Pilih algoritma string matching yang digunakan: ")
        fmt.Scanln(&algo)

        if (algo == 1) {
            result_uji = stringmatching.Boyermoore(string(content), seq_check)
        } else if (algo == 2) {
            result_uji = stringmatching.KMP(string(content), seq_check)
        } 

        if (result_uji == 1) {
            sim = 100
        } else {
            sim = float64(stringmatching.LCS(string(content), seq_check)) / float64(len(seq_check))
            if (sim >= 80) {
                result_uji = 1
            }
        }
        insert, err := db.Query("INSERT INTO data_uji (tanggal_tes, nama_pengguna, nama_penyakit, similarity, status_tes) VALUES ( ?, ? , ?, ?, ? )", time.Now().Format("2006-01-02"), pengguna, penyakit, sim, result_uji)
        if err != nil {
            log.Printf("Error %s when insert to DB\n", err)
        }
        defer insert.Close()
    }
}