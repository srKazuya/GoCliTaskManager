package utils

import (
	"clitask/models"
	"encoding/json"
	"fmt"
	"os"
)

func Unparse() []models.Task {
	fileData, err := os.ReadFile("tasks.json")

	var tasks []models.Task
	if err == nil && len(fileData) > 0 {

		err = json.Unmarshal(fileData, &tasks)
		if err != nil {
			fmt.Println("Ошибка при чтении задач", err)
			os.Exit(1)
		}
	}
	return tasks
}

func Parse(t []models.Task) {
	data, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		os.Exit(1)
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка сериализацции")
	}
	fmt.Println("Задачи обновлены")
}
