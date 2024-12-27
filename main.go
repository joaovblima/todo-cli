package main

import "fmt"

func main() {
	fmt.Println("Bem vindo ao TODO LIST CLI")
	mostrarMenu()
}

func mostrarMenu(){
	fmt.Println("\nEscolha uma opção:")
	fmt.Println("1. Adicionar Tarefa")
	fmt.Println("2. Listar Tarefas")
	fmt.Println("3. Concluir Tarefa")
	fmt.Println("4. Remover Tarefa")
	fmt.Println("5. Sair")
}