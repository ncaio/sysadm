//
//      Ncaio
//      caiogore _|_ gmail _|_ com
//
package main

//
//
//
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "os"
	"os/exec"
)

//
//
//
type Todo struct {
	Resultado bool   `json:"resultado"`
	Url       string `json:"url"`
	Token     string `json:"token"`
	Telefone  string `json:"telefone"`
}

type Todos []Todo

//
//
//
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/sms/{smsID}/{smsNum}", Sms).Methods("GET")
	router.HandleFunc("/url/{url}/services/email/confirmar/{chave}/{smsNum}", Smsurl).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//
//
//
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SET-SMS-GATEWAY - Servico de envio de SMS")
}

//
//
//
func Sms(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	smsID := vars["smsID"]
	smsNum := vars["smsNum"]
	if len(smsID) > 0 {
		todos := Todos{
			Todo{
				Resultado: true,
				Token:     smsID,
				Telefone:  smsNum,
			},
		}
		if err := json.NewEncoder(w).Encode(todos); err != nil {
			log.Println(err)
		}
	}
	fsms(smsID, smsNum)
}

//
//
//
func Smsurl(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	smsNum := vars["smsNum"]
	chave := vars["chave"]
	url := vars["url"]
	if len(url) > 0 {
		todos := Todos{
			Todo{
				Resultado: true,
				Url:       url,
				Token:     chave,
				Telefone:  smsNum,
			},
		}
		if err := json.NewEncoder(w).Encode(todos); err != nil {
			log.Println(err)
		}
	}
	furl(url, chave, smsNum)
}

//
//
//

func fsms(sid string, snum string) {
	cmd := exec.Command("sh", "-c", "echo \"TOKEN para validar: "+sid+"\" | /usr/bin/gnokii --sendsms "+snum)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(stdout))
	}
	//log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	//out, _ := os.Create("sms-gateway.log")
	//prefix := "SMS GATEWAY API " + sid + " " + snum + " "
	//flag := log.LstdFlags | log.Lmicroseconds | log.Lshortfile
	//newLog := log.New(out, prefix, flag)
	//newLog.Println("")
}

func furl(url string, chave string, snum string) {
	cmd := exec.Command("sh", "-c", "echo \"Para recuperar sua senha, acesse o link: https://"+url+"/services/email/confirmar/"+chave+"\" | /usr/bin/gnokii --sendsms "+snum)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(stdout))
	}
	//log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	//out, _ := os.Create("sms-gateway.log")
	//prefix := "SMS GATEWAY API " + sid + " " + snum + " "
	//flag := log.LstdFlags | log.Lmicroseconds | log.Lshortfile
	//newLog := log.New(out, prefix, flag)
	//newLog.Println("")
}
