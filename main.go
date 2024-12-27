package main

import "fmt"

type Tarefas struct {
	ID int
	Descrição string
	Concluida bool
}

var tarefas []Tarefas
var nextId int = 1

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

func adicionarTarefas(descricao string) {
	tarefa := Tarefas {
		ID: nextId,
		Descrição: descricao,
		Concluida: false,
	}

	tarefas = append(tarefas, tarefa)
	nextId++
	fmt.Println("Tarefa criada.")
}

