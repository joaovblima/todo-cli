package helper

import "fmt"


type Tasks struct {
	ID int
	Descrição string
	Concluida bool
}

var tarefas = []Tasks{}
var nextId int = 1


func AdicionarTarefas(descricao string) {
	tarefa := Tasks {
		ID: nextId,
		Descrição: descricao,
		Concluida: false,
	}

	tarefas = append(tarefas, tarefa)
	nextId++
	fmt.Println("Tarefa criada.")
	fmt.Println()
}

func ListarTarefas() {
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

func CompletarTarefa(id int) {
	for i, tarefa := range tarefas {
		if tarefa.ID == id {
			tarefas[i].Concluida = true
			fmt.Println("Tarefa marcada como concluida. ")
			return
		}
	}
	fmt.Println("Tarefa não encontrada")
}




func ExcluirTarefa(id int) {
	for i, tarefa := range tarefas {
		if tarefa.ID == id {
			tarefas = append(tarefas[:i], tarefas[i + 1:]... )
			fmt.Println("Tarefa removida com sucesso!")
			return 
		}
	}
	fmt.Println("Tarefa nao encontrada.")
}