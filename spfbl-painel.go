//
//	Noilson Caio - caiogore gmail com
//
package main
//
//
//
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"net/http"
)

//	Readlines From File	
//	https://siongui.github.io/2016/04/06/go-readlines-from-file-or-string/#readlines-from-file
//

func File2lines(filePath string) []string {
      f, err := os.Open(filePath)
      if err != nil {
              panic(err)
      }
      defer f.Close()

      var lines []string
      scanner := bufio.NewScanner(f)
      for scanner.Scan() {
			lines = append(lines, scanner.Text())
      }
      if err := scanner.Err(); err != nil {
              fmt.Fprintln(os.Stderr, err)
      }

      return lines
}
//
//
//
func handler(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
//
//
//
	var saida []string
	saida = File2lines("/go/painelSPFBL/oficina/spfbl.2017-04-05.log")
	fmt.Fprintf(w, "<div align=center><h1>[- PAINEL DE CONTROLE CLIENTE -]</h1> </div>")
//
//	LOOP
//
	for i := range saida {
	if strings.Contains(saida[i], "yourhostnamehere") {
	field := strings.Fields(saida[i])
	date := (strings.Replace(field[0], "T", " hora: ", -1))
	//if strings.Contains(field[14], "spfbl.set.rn.gov.br") {
//
//	PASS
//
	if strings.Contains(field[13], "PASS") && strings.Contains(field[14], "http") {
		fmt.Fprintf(w, "<div style=background-color:white align=left><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong><a href=%s target=_blank>%s</a></strong></font></p></div>",date,field[9],field[10],field[14],field[13])
//
//	BLOCK
//
	} else if strings.Contains(field[13], "BLOCK") && strings.Contains(field[14], "http") {
		fmt.Fprintf(w, "<div style=background-color:red align=left><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong><a href=%s target=_blank>%s</a></strong></font></p></div>",date,field[9],field[10],field[14],field[13])
//
//	WHITE	
//
	} else if strings.Contains(field[13], "WHITE") && strings.Contains(field[14], "http"){
		fmt.Fprintf(w, "<div style=background-color:white align=left><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong><a href=%s target=_blank>%s</a></strong></font></p></div>",date,field[9],field[10],field[14],field[13])
//
//	SOFTFAIL
//
	} else if strings.Contains(field[13], "SOFTAIL") && strings.Contains(field[14], "http"){
                fmt.Fprintf(w, "<div style=background-color:white align=left><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong><a href=%s target=_blank>%s</a></strong></font></p></div>",date,field[9],field[10],field[14],field[13])
//
//	NEUTRAL
//
        } else if strings.Contains(field[13], "NEUTRAL") && strings.Contains(field[14], "http"){
                fmt.Fprintf(w, "<div style=background-color:white align=left><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong><a href=%s target=_blank>%s</a></strong></font></p></div>",date,field[9],field[10],field[14],field[13])
//
//	NONE
//
        } else if strings.Contains(field[13], "NONE") && strings.Contains(field[14], "http"){
                fmt.Fprintf(w, "<div style=background-color:white align=left><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong><a href=%s target=_blank>%s</a></strong></font></p></div>",date,field[9],field[10],field[14],field[13])
//
//	FLAG
//
        } else if strings.Contains(field[13], "FLAG") && strings.Contains(field[14], "http"){
                fmt.Fprintf(w, "<div style=background-color:yellow align=left><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong><a href=%s target=_blank>%s</a></strong></font></p></div>",date,field[9],field[10],field[14],field[13])
//
//	OTHERS
//
	} //else {
	//fmt.Fprintf(w, "<div style=background-color:cyan align=left><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong>%s</strong></font></p></div>",date,field[9],field[10],field[13])
	//}
	//fmt.Fprintf(w,"<hr>")
	}
	}
}
//
//	FUNC MAIN
//
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
