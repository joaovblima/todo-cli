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
	for {
		fmt.Println("Bem vindo ao TODO LIST CLI")
		mostrarMenu()

		var escolha int
		fmt.Scanln(&escolha)

		switch escolha {
		case 1:
			fmt.Println("Descreva a tarefa: ")
			var descricao string
			fmt.Scanln(&descricao)
			adicionarTarefas(descricao)

		case 2:
			listarTarefas()
		
		case 3:
			fmt.Println("Digite o id para completar a tarefa: ")
			var id int 
			fmt.Scanln(&id)
			completarTarefa(id)

		case 4: 
			fmt.Println("Digite o id para excluir tarefa: ")
			var id int 
			fmt.Scanln(&id)
			excluirTarefa(id)
		
		case 5:
			fmt.Println("Saindo...")
			return
		
		default:
			fmt.Println("Opção invalida... Tente novamente")
		}

	}

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

func excluirTarefa(id int) {
	for i, tarefa := range tarefas {
		if tarefa.ID == id {
			tarefas = append(tarefas[:i], tarefas[i + 1:]... )
			fmt.Println("Tarefa removida com sucesso!")
			return
		}
	}
	fmt.Println("Tarefa nao encontrada.")
}