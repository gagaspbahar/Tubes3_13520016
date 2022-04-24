package add

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
)

func CheckSequence(input string) bool {
    res1, _ := regexp.MatchString(`^[CAGT]+$`, input)
    return res1
}

func AddPenyakit(db *sql.DB) {
    var filename string
    var penyakit string
    fmt.Print("Masukan nama file: ")
    fmt.Scanln(&filename)
    fmt.Print("Masukan nama penyakit: ")
    fmt.Scanln(&penyakit)
    content, err := os.ReadFile("data/sequence/" + filename)
    if err != nil {
        log.Fatal(err)
    }
    result := CheckSequence(string(content))
    if (result) {
        insert, err := db.Query("INSERT INTO sequence_penyakit VALUES ( ?, ? )", penyakit, content)
        if err != nil {
            log.Printf("Error %s when insert to DB\n", err)
        }
        defer insert.Close()
    }
}