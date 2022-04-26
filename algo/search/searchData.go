package search_data

import (
	"database/sql"
	"fmt"
    "bufio"
	"os"
	"regexp"
    "strings"
    "strconv"
    "log"
)

type Record struct {
    tanggal string
    nama_pengguna string
    nama_penyakit string
    similarity float64
    status_tes int
}

func CheckSequence(input string) (string, string) {
    var ans, ans2 string
    jan := "[jJ][aA][nN]([uU][aA][rR][yY])?"
    feb := "[fF][eE][bB]([rR][uU][aA][rR][yY])?"
    mar := "[mM][aA][rR]([cC][hH])?"
    apr := "[aA][pP][rR]([iI][lL])?"
    may := "[mM][aA][yY]"
    jun := "[jJ][uU][nN]([eE])?"
    jul := "[jJ][uU][lL]([yY])?"
    aug := "[aA][uU][gG]([uU][sS][tT])?"
    sep := "[sS][eE][pP]([tT][eE][mM][bB][eE][rR])?"
    oct := "[oO][cC][tT]([oO][bB][eE][rR])?"
    nov := "[nN][oO][vV]([eE][mM][bB][eE][rR])?"
    dec := "[dD][eE][cC]([eE][mM][bB][eE][rR])?"
    day_28 := "((0[1-9]|1[0-9]|2[0-8])(\\s)*(-|\\/| )(\\s)*(0[1-9]|1[012]|" + jan + "|" + feb + "|" + mar + "|" + apr + "|" + may + "|" + jun + "|" + jul + "|" + aug + "|" + sep + "|" + oct + "|" + nov + "|" + dec + "))"
    day_30 := "((29|30)(\\s)*(-|\\/| )(\\s)*(0[13456789]|1[012]|" + jan + "|" + mar + "|" + apr + "|" + may + "|" + jun + "|" + jul + "|" + aug + "|" + sep + "|" + oct + "|" + nov + "|" + dec + "))"
    day_31 := "(31(\\s)*(-|\\/| )(\\s)*(0[13578]|1[02]|" + jan + "|" + mar + "|" + may + "|" + jul + "|" + aug + "|" + oct + "|" + dec + "))"
    day_month := "(" + day_28 + "|" + day_30 + "|" + day_31 + ")"
    year := "([0-1]\\d{3}|20[0-1][0-9]|202[0-2])"
    leap :=  "29(\\s)*(-|\\/| )(\\s)*(02|" + feb + ")(\\s)*(-|\\/| )(\\s)*([0-1][0-9]|20)(00|04|08|12|16|20|24|28|32|36|40|44|48|52|56|60|64|68|72|76|80|84|88|92|96)"
    penyakit := "((\\s)+([a-zA-Z])+)*(\\s)*"
    regex := "(^(\\s)*" + day_month + "(\\s)*(-|\\/| )(\\s)*" + year + penyakit + "$)|(^(\\s)*" + leap + penyakit + "$)|(^(\\s)*(([a-zA-Z])+(\\s)*)+$)"
    r , _ := regexp.Compile(regex)
    r_date, _ := regexp.Compile("(^(\\s)*" + day_month + "(\\s)*(-|\\/| )(\\s)*" + year + ")|(^" + leap + ")")
    res := r.MatchString(input)
    if (res) {
        ans = r_date.FindString(input)
        ans2 = strings.ReplaceAll(input, ans, "")
    }
    return ans, ans2
}

func Clear_whitespace(date string, penyakit string) (string, string) {
    var ans_date string
    var l = []string{"jan", "feb", "mar", "apr","may","jun","jul","aug","sep","oct","nov","dec"}
    date = strings.ReplaceAll(date, " ", "")
    date = strings.ReplaceAll(date, "/", "")
    date = strings.ReplaceAll(date, "-", "")
    if (date != "") {
        a := []rune(date)
        day := string(a[0:2])
        month := string(a[2:len(a)-4])
        year := string(a[len(a)-4:])
        if len(month) != 2 {
            for i := 0; i < 12; i++ {
                if (strings.ToLower(month[0:3]) == l[i]) {
                    if (i < 9) {
                        month = "0" + strconv.Itoa(i+1)
                    } else {
                        month = strconv.Itoa(i+1)
                    }
                    break
                }
            }
        }
        ans_date = year + "-" + month + "-" + day
    }
    penyakit = strings.TrimSpace(penyakit)
    space := regexp.MustCompile(`\s+`)
    penyakit = space.ReplaceAllString(penyakit, " ")
    return ans_date, penyakit
}

func Search_db(db *sql.DB) {
    var results *sql.Rows
    var err error
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Masukkan tanggal (day month year) dan nama penyakit: ")
    data, _ := reader.ReadString('\n')
	data = strings.TrimSuffix(data, "\n")
    data = strings.TrimSuffix(data, "\r")
    date, penyakit := CheckSequence(data)
    date, penyakit = Clear_whitespace(date, penyakit)
    if (date != "" && penyakit != "") {
        results, err = db.Query("SELECT tanggal_tes, nama_pengguna, nama_penyakit, similarity, status_tes FROM data_uji WHERE tanggal_tes = ? AND nama_penyakit = ?",date, penyakit)
    } else if (penyakit != "") {
        results, err = db.Query("SELECT tanggal_tes, nama_pengguna, nama_penyakit, similarity, status_tes FROM data_uji WHERE nama_penyakit = ?", penyakit)
    } else if (date != "") {
        results, err = db.Query("SELECT tanggal_tes, nama_pengguna, nama_penyakit, similarity, status_tes FROM data_uji WHERE tanggal_tes = ?",date)
    } else {
        return
    }
    if err != nil {
        log.Printf("Error %s when SELECT from DB\n", err)
    }
    for results.Next() {
        var record Record 
        // for each row, scan the result into our tag composite object
        err = results.Scan(&record.tanggal, &record.nama_pengguna, &record.nama_penyakit, &record.similarity, &record.status_tes)
        if err != nil {
            log.Printf("Error %s when insert to DB\n", err)
        }
        fmt.Println(record)
    }
}