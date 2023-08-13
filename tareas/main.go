package main

import "fmt"

type Task struct {
	title       string
	description string
	complete    bool
}

type ListTask struct {
	tasks []Task
}

// metodo agregar tareas
func (listTasks *ListTask) addTarea(task Task) {
	listTasks.tasks = append(listTasks.tasks, task)
}

// metodo para completar tarea
func (listTasks *ListTask) completeTaks(index int) {
	listTasks.tasks[index].complete = true
}

func (listTasks *ListTask) editTask(index int, task Task) {
	listTasks.tasks[index] = task
}

// eliminar tarea
func (listTasks *ListTask) deleteTask(index int) {
	listTasks.tasks = append(listTasks.tasks[:index], listTasks.tasks[index+1:]...)
}

func main() {
	//Instancia de lista de tareas
	listTasks := ListTask{}

	for{
		
	}

	fmt.Println("Lista de tareas")
	fmt.Println("===============================")
	for index, task := range listTasks.tasks {
		fmt.Printf("%d. %s - %s - completado %t \n", index, task.title, task.description, task.complete)
	}
	fmt.Println("===============================")
}
