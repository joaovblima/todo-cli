package main

import "fmt"

type Tasks struct {
	ID int
	Descrição string
	Concluida bool
}

var tarefas []Tasks
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
	tarefa := Tasks {
		ID: nextId,
		Descrição: descricao,
		Concluida: false,
	}

	tarefas = append(tarefas, tarefa)
	nextId++
	fmt.Println("Tarefa criada.")
}

func listarTarefas() {
	if len(tarefas) == 0{
		fmt.Println("Vocẽ ainda não possui tarefas adicionadas")
		return
	}

	for _, tarefa := range tarefas {
		status := "Pendente"
		
		if tarefa.Concluida {
			status = "Concluida"
		}
		fmt.Printf("ID: %d | Descrição: %s | Status: %s\n", tarefa.ID, tarefa.Descrição, status)
	}
}

func completarTarefa(id int) {
	for i, tarefa := range tarefas {
		if tarefa.ID == id {
			tarefas[i].Concluida = true
			fmt.Println("Tarefa marcada como concluida. ")
			return
		}
	}
	fmt.Println("Tarefa não encontrada")
}