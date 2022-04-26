package add

import (
	"database/sql"
	"log"
	"os"
	"regexp"
)

func CheckSequence(input string) bool {
	res1, _ := regexp.MatchString(`^[CAGT]+$`, input)
	return res1
}

func AddPenyakit(db *sql.DB, filename, penyakit string) {
	content, err := os.ReadFile("data/sequence/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	result := CheckSequence(string(content))
	if result {
		insert, err := db.Query("INSERT INTO sequence_penyakit VALUES ( ?, ? )", penyakit, content)
		if err != nil {
			log.Printf("Error %s when insert to DB\n", err)
		}
		defer insert.Close()
	}
}
