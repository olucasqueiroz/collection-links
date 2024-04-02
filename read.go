package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Link struct {
	Nome string `json:"nome"`
	Link string `json:"link"`
}

type LinksData struct {
	Links []Link `json:"dados"`
}

func main() {
	// Verificar se o arquivo existe
	_, err := os.Stat("links.json")

	if err != nil {
		fmt.Println("O arquivo links.json não existe.")
		return
	}

	// Ler as informações do arquivo
	file, err := ioutil.ReadFile("links.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo links.json:", err)
		return
	}

	var data LinksData
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Erro ao decodificar o arquivo links.json:", err)
		return
	}

	// Imprimir as informações dos links
	fmt.Println("Informações dos Links:")
	for _, link := range data.Links {
		fmt.Println("Nome:", link.Nome)
		fmt.Println("Link:", link.Link)
		fmt.Println("010101010101010101010101010")
	}
}
