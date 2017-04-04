package main

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
	saida = File2lines("/go/painelSPFBL/oficina/spfbl.log")
	fmt.Fprintf(w, "<div align=center><h1>PAINEL DE CONTROLE</h1> </div>")
	for i := range saida {
	remetente := strings.Fields(saida[i])
	date := (strings.Replace(remetente[0], "T", " hora: ", -1))
	if strings.Contains(remetente[13], "PASS") {
		fmt.Fprintf(w, "<div style=background-color:yellow align=center><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong>%s</strong></font></p></div>",date,remetente[9],remetente[10],remetente[13])
	} else if strings.Contains(remetente[13], "BLOCK") {
		fmt.Fprintf(w, "<div style=background-color:red align=center><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong>%s</strong></font></p></div>",date,remetente[9],remetente[10],remetente[13])
	} else if strings.Contains(remetente[13], "GREYLIST"){
		fmt.Fprintf(w, "<div style=background-color:grey align=center><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong>%s</strong></font></p></div>",date,remetente[9],remetente[10],remetente[13])
	} else { fmt.Fprintf(w, "<div style=background-color:cyan align=center><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong>%s</strong></font></p></div>",date,remetente[9],remetente[10],remetente[13])
	}
	}
}
//
//
//
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
root@c45f6402a2d9:/go/painelSPFBL/oficina# cat leia.go 
package main

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
	saida = File2lines("/go/painelSPFBL/oficina/spfbl.log")
	fmt.Fprintf(w, "<div align=center><h1>PAINEL DE CONTROLE</h1> </div>")
	for i := range saida {
	remetente := strings.Fields(saida[i])
	date := (strings.Replace(remetente[0], "T", " hora: ", -1))
	if strings.Contains(remetente[13], "PASS") {
		fmt.Fprintf(w, "<div style=background-color:yellow align=center><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong>%s</strong></font></p></div>",date,remetente[9],remetente[10],remetente[13])
	} else if strings.Contains(remetente[13], "BLOCK") {
		fmt.Fprintf(w, "<div style=background-color:red align=center><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong>%s</strong></font></p></div>",date,remetente[9],remetente[10],remetente[13])
	} else if strings.Contains(remetente[13], "GREYLIST"){
		fmt.Fprintf(w, "<div style=background-color:grey align=center><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong>%s</strong></font></p></div>",date,remetente[9],remetente[10],remetente[13])
	} else { fmt.Fprintf(w, "<div style=background-color:cyan align=center><p><font size=3 face=arial>Em <strong>%s</strong> o remetente: <strong>%s</strong> MTA: <strong>%s</strong> foi detectado como: <strong>%s</strong></font></p></div>",date,remetente[9],remetente[10],remetente[13])
	}
	}
}
//
//
//
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
