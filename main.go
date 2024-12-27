package main

import ("fmt" 
		"todo/helper"
)


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
			helper.AdicionarTarefas(descricao)
		case 2:
			helper.ListarTarefas()
		
		case 3:
			fmt.Println("Digite o id para completar a tarefa: ")
			var id int 
			fmt.Scanln(&id)
			helper.CompletarTarefa(id)

		case 4: 
			fmt.Println("Digite o id para excluir tarefa: ")
			var id int 
			fmt.Scanln(&id)
			helper.ExcluirTarefa(id)
		
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
