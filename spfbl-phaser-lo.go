package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
//
//
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

func main() {
	var saida []string
	saida = File2lines("/go/painelSPFBL/oficina/spfbl.log")
	for i := range saida {
	remetente := strings.Fields(saida[i])
	date := (strings.Replace(remetente[0], "T", " hora: ", -1))
	fmt.Println("Em: " + date, "- O remetente: " + remetente[9], "- MTA: " +remetente[10],"Foi detectado como: " + remetente[13])
	}
}
