package main

import "fmt"

func main(){

	nome := "Ana"
	idade := 30
	altura := 1.68
	var pontos int
	const empresa = "AtendePro"

	media := float64(idade) / 7

	fmt.Println(nome, idade, altura, pontos, empresa)
	fmt.Println(media)

}