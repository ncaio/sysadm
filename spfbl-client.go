//
//
//
package main

//
//
//
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"github.com/pelletier/go-toml"
)

//
//
//
func main() {
	//
	//
	//
	leitor := bufio.NewReader(os.Stdin)
	fmt.Print("White list add: ")
	consulta,_ := leitor.ReadString('\n')
	status := (white(consulta,"BLOCK","DROP"))
	fmt.Println(status)
}
//
//
//
func white(email string, comando string, acao string) (resultado string){
        //
        //
        //
        config, err := toml.LoadFile("/etc/painel-spfbl.toml")
        if err != nil {
                fmt.Println("Error: ", err.Error())
        }
        host := config.Get("server.host").(string)
        port := config.Get("server.port").(string)
        //
        //
        //
	com, _ := net.Dial("tcp", host + ":" + port)
	defer com.Close()
	//
	//
	//
	if comando == "WHITE" {
		if acao == "ADD"{
			fmt.Fprintf(com,"WHITE SENDER " + email + "\n")
			result, _ := bufio.NewReader(com).ReadString('\n')
			return result
		}
		if acao == "DROP"{
			fmt.Fprintf(com,"WHITE DROP " + email + "\n")
			result, _ := bufio.NewReader(com).ReadString('\n')
			return result
		}
	}
	if comando == "BLOCK" {
		if acao == "ADD"{
			fmt.Fprintf(com,"BLOCK ADD " + email + "\n")
			result, _ := bufio.NewReader(com).ReadString('\n')
			return result
		}
		if acao == "DROP"{
                        fmt.Fprintf(com,"BLOCK DROP " + email + "\n")
                        result, _ := bufio.NewReader(com).ReadString('\n')
                        return result
                }
	}
	return
}
//
//
//
