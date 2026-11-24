package main

import "fmt"
import "time"

func main(){
	fmt.Println("Olá, Go!" + time.Now().Format(time.DateTime))
}