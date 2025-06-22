package commands

import (
	"bufio"
	"clitask/models"
	"clitask/utils"
	"fmt"
	"strings"
	"time"
)

type Command interface {
	Execute()
}

type AddCommand struct {
	Reader *bufio.Reader
}

type DeleteCommand struct {
	Reader *bufio.Reader
}

type UpdateCommand struct {
	Reader *bufio.Reader
}

type ListCommand struct {
	Reader *bufio.Reader
}

type ListNotDoneCommand struct {
	ListCommand
}

type MarkAsDoneCommand struct {
	// Reader *bufio.Reader
	ListNotDoneCommand
}

func (c *AddCommand) Execute() {
	fmt.Print("Введите задачу: ")
	desc, _ := c.Reader.ReadString('\n')
	desc = strings.TrimSpace(desc)
	task := models.NewTask(desc)

	tasks := utils.Unparse()

	if len(tasks) > 0 {
		task.Id = tasks[len(tasks)-1].Id + 1
	}
	tasks = append(tasks, task)

	utils.Parse(tasks)

	fmt.Printf("Задача добавлена: %q\n", task.Description)
}

func (c *DeleteCommand) Execute() {
	tasks := utils.Unparse()

	idToIndex := utils.BuildIdToIndexMap(tasks)

	for {
		num, err := utils.GetTaskId(c.Reader)
		if err != nil {
			fmt.Println(err)
			continue
		}
		index, ok := idToIndex[num]
		if !ok {
			fmt.Println("Задача с таким ID не найдена")
			continue
		}

		tasks = append(tasks[:index], tasks[index+1:]...)
		fmt.Println("Задача удалена")

		utils.Parse(tasks)
		break
	}
}

func BuildIdToIndexMap(tasks []models.Task) any {
	panic("unimplemented")
}
func (c *UpdateCommand) Execute() {
	tasks := utils.Unparse()
	idToIndex := utils.BuildIdToIndexMap(tasks)

	for {
		taskId, err := utils.GetTaskId(c.Reader)
		if err != nil {
			fmt.Println(err)
			continue
		}

		index, ok := idToIndex[taskId]
		if !ok {
			fmt.Println("Задача с таким ID не найдена")
			continue
		}
		task := &tasks[index]
		for {
			fmt.Printf("Выбрана задача %d: %q\nВведите новое описание: ", task.Id, task.Description)
			input, err := c.Reader.ReadString('\n')
			if err != nil {
				fmt.Println("Введите новое описание для команды ", err)
				continue
			}
			input = strings.TrimSpace(input)
			task.Description = input
			task.UpdatedAt = time.Now()
			break
		}
		utils.Parse(tasks)
		break
	}

}

func (c *ListCommand) Execute() {
	tasks := utils.Unparse()
	for _, task := range tasks {
		c.FormatOutput(task)
	}
}

func (c *ListCommand) FormatOutput(task models.Task) {
	fmt.Printf(
		" \033[1;34m %d \033[1;34mЗадача:\033[0m %q, \033[1;36mСтатус:\033[0m %s \033[1;32mСоздана:\033[0m %s, \033[1;33mРедактирована:\033[0m %s\n",
		task.Id,
		task.Description,
		task.Status,
		task.CreatedAt.Format("Mon, 02.01.2006 15:04"),
		task.UpdatedAt.Format("Mon, 02.01.2006 15:04"),
	)
}

func (c *ListNotDoneCommand) Execute() {
	tasks := utils.Unparse()
	for _, task := range tasks {
		if task.Status == "В работе" {
			c.FormatOutput(task)
		} else {
			continue
		}
	}
}

func (c *MarkAsDoneCommand) Execute() {
	c.ListNotDoneCommand.Execute()
	tasks := utils.Unparse()

	idToIndex := utils.BuildIdToIndexMap(tasks)

	for {
		taskId, err := utils.GetTaskId(c.Reader)
		if err != nil {
			fmt.Println(err)
			continue
		}

		index, ok := idToIndex[taskId]
		if !ok {
			fmt.Println("Задачи с таким ID  нет")
			break
		}

		task := &tasks[index]

		if task.Status == "В работе" {
			task.Status = "Выполнена"
			task.UpdatedAt = time.Now()
			fmt.Printf("Статус задачи %q, изменен на: %q\n", task.Description, task.Status)
		} else {
			fmt.Printf("Задача %q уже имеет статус: %q\n", task.Description, task.Status)
		}
		break
	}

	utils.Parse(tasks)
}
