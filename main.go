package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	complete    bool
}

type TaskList struct {
	task []Task
}

// Add task
func (l *TaskList) addTask(t Task) {
	l.task = append(l.task, t)
}

// Task complete
func (l *TaskList) taskComplete(index int) {
	l.task[index].complete = true
}

// Edit task
func (l *TaskList) editTask(index int, t Task) {
	l.task[index] = t
}

// Remove task
func (l *TaskList) removeTask(index int) {
	l.task = append(l.task[:index], l.task[index+1:]...)
}

func main() {
	//Instance taskList
	list := TaskList{}
	read := bufio.NewReader(os.Stdin)

	for {
		var option int
		fmt.Println(
			"Select option:\n",
			"1. Add task:\n",
			"2. Complete task:\n",
			"3. Edit task:\n",
			"4. Remove task:\n",
			"5. Exit:",
		)

		fmt.Print("Type the option \n")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var t Task
			fmt.Print("Type task name \n")
			t.name, _ = read.ReadString('\n')
			fmt.Print("Type task description \n")
			t.description, _ = read.ReadString('\n')
			list.addTask(t)
			fmt.Print("Task added \n")
		case 2:
			var index int
			fmt.Print("Type index of the task that you want to complete \n")
			fmt.Scanln(&index)
			list.taskComplete(index)
			fmt.Println("Task complete correctly")
		case 3:
			var index int
			var t Task
			fmt.Print("Type the index of the task that you want to edit \n")
			fmt.Scanln(&index)
			fmt.Print("Type the task name \n")
			t.name, _ = read.ReadString('\n')
			fmt.Print("Type the task description \n")
			t.description, _ = read.ReadString('\n')
			list.editTask(index, t)
			fmt.Println("Task updated correctly")
		case 4:
			var index int
			fmt.Print("Type the index of the task that you want to remove \n")
			fmt.Scanln(&index)
			list.removeTask(index)
			fmt.Println("Task removed correctly")
		case 5:
			fmt.Print("Exit... \n")
			return
		default:
			fmt.Print("Invalid option \n")
		}

		fmt.Println("Task list:")
		for i, t := range list.task {
			fmt.Printf("%d. %s - %s - Complete: %t \n", i, t.name, t.description, t.complete)
		}
	}
}
