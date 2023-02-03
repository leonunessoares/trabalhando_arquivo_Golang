
package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"strconv"
)
var agora = time.Now()
var path = "teste" + "_" + strconv.Itoa(agora.Hour()) + "_" + strconv.Itoa(agora.Minute()) +  ".txt"

func main () {
	criar()
	abrir_escrever()
	ler()
}
	
func criar () {
    //Verifica se o arquivo existe
    var _, err = os.Stat(path)

    //Cria o arquivo
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if err != nil {
            log.Fatal(err)
        } 
        defer file.Close()
    }
}

 func abrir_escrever (){
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if err != nil {
        log.Fatal(err)
    }
	
	_, err = file.WriteString("Primeira linha - Golang\n")
	_, err = file.WriteString("Segunda linha - Golang\n")
    if err != nil {
        log.Fatal(err)
    }

	defer file.Close()

 } 
  
  func ler (){
	content, err := os.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(content))

  }