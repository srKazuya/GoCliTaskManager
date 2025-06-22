package main

import (
	"bufio"
	"clitask/commands"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := map[string]commands.Command{
		"добавить":        &commands.AddCommand{Reader: reader},
		"удалить":         &commands.DeleteCommand{Reader: reader},
		"обновить":        &commands.UpdateCommand{Reader: reader},
		"показать задачи": &commands.ListCommand{Reader: reader},
		"показать невыполненные": &commands.ListNotDoneCommand{
			ListCommand: commands.ListCommand{
				Reader: reader,
			},
		},
		"пометить как выполненную" : &commands.MarkAsDoneCommand{
			ListNotDoneCommand: commands.ListNotDoneCommand{
				ListCommand: commands.ListCommand{
					Reader: reader,
				},
			},
		},
	}

	for {
		fmt.Print("\033[1;34mВыберите команду:  \033[1;32m\nдобавить\nобновить\nудалить\nпоказать задачи\nпоказать невыполненные\nпометить как выполненную\n")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		cmd, ok := commands[input]
		if !ok {
			fmt.Println("Неизвестная команда")
			continue
		}
		cmd.Execute()

	}
}
