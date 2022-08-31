package main

import "fmt"

//lista de tareas
type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appenTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}
func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

type Task struct {
	name         string
	descripction string
	complete     bool
}

func (t *Task) imprimir() {
	fmt.Printf("Nombre: %s\n Descripción %s, \n Completado: %t\n", t.name, t.descripction, t.complete)
}
func (t *Task) taskComplete() {
	t.complete = true
}
func main() {
	t1 := Task{"Curso de Go", "Completar curso de Go este mes", false}
	t2 := Task{"Curso de HTML", "Curso de HTML", true}
	fmt.Println("Tarea 1", t1)
	fmt.Println("Tarea 2", t2)

	lista := TaskList{}
	lista.appenTask(&t1)
	lista.appenTask(&t2)
	fmt.Println(lista)

	t3 := Task{"Curso de Carbon", "Curso de Carbon", false}
	lista.appenTask(&t3)
	fmt.Println(lista)

	for i, task := range lista.tasks {
		fmt.Println(i, task.name)
	}

}
