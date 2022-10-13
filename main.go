package main

import (
	"blog/pkg/configs"
	"blog/pkg/services"
	"fmt"
)

func main() {
	// Carrega as configurações da API
	configs.Load()
	PORT := configs.GetAPIPort()

	// Inicia um novo servidor
	srv := services.Server()
	srv.Addr = PORT

	err := srv.ListenAndServe()

	if err != nil {
		fmt.Printf("Erro ao subir servidor: %v", err)
	}
}
