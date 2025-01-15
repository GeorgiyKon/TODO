package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks []string

// Функция для отображения текущих задач
func displayTasks() {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст.")
	} else {
		fmt.Println("Список задач:")
		for i, task := range tasks {
			fmt.Printf("%d: %s\n", i+1, task)
		}
	}
}

// Функция для добавления задачи
func addTask(task string) {
	tasks = append(tasks, task)
	fmt.Println("Задача добавлена:", task)
}

// Функция для удаления задачи
func deleteTask(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Некорректный индекс.")
		return
	}
	fmt.Println("Задача удалена:", tasks[index])
	tasks = append(tasks[:index], tasks[index+1:]...)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите команду (add, delete, list, quit): ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)

		switch parts[0] {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Используйте: add <задача>")
				continue
			}
			task := strings.Join(parts[1:], " ")
			addTask(task)

		case "delete":
			if len(parts) != 2 {
				fmt.Println("Используйте: delete <номер задачи>")
				continue
			}
			var index int
			_, err := fmt.Sscan(parts[1], &index)
			if err != nil || index < 1 || index > len(tasks) {
				fmt.Println("Некорректный индекс.")
				continue
			}
			deleteTask(index - 1)

		case "list":
			displayTasks()

		case "quit":
			fmt.Println("Выход из приложения.")
			return

		default:
			fmt.Println("Неизвестная команда. Попробуйте: add, delete, list, quit.")
		}
	}
}
