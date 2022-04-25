package search_data

import (
	"database/sql"
	"fmt"
    "bufio"
	"os"
	"regexp"
    "strings"
)

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
    leap :=  "29(\\s)*(-|\\/| )(\\s)*(02|FEB(RUARY)?)(\\s)*(-|\\/| )(\\s)*([0-1][0-9]|20)(00|04|08|12|16|20|24|28|32|36|40|44|48|52|56|60|64|68|72|76|80|84|88|92|96)"
    penyakit := "((\\s)+([a-zA-Z])+)*(\\s)*"
    regex := "(^(\\s)*" + day_month + "(\\s)*(-|\\/| )(\\s)*" + year + penyakit + "$)|(^" + leap + penyakit + "$)|(^(([a-zA-Z])+(\\s)*)+$)"
    r , _ := regexp.Compile(regex)
    r_date, _ := regexp.Compile("(^(\\s)*" + day_month + "(\\s)*(-|\\/| )(\\s)*" + year + ")|(^" + leap + ")")
    res := r.MatchString(input)
    if (res) {
        ans = r_date.FindString(input)
        ans2 = strings.ReplaceAll(input, ans, "")
    }
    return ans, ans2
}

func Search_db(db *sql.DB) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Masukkan tanggal (day month year) dan nama penyakit: ")
    data, _ := reader.ReadString('\n')
	data = strings.TrimSuffix(data, "\n")
    data = strings.TrimSuffix(data, "\r")
    date, penyakit := CheckSequence(data)
    fmt.Println(date)
    fmt.Println(penyakit)
}