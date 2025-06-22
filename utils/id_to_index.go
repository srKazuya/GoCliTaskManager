package utils

import (
	"clitask/models"
	"fmt"
)

func BuildIdToIndexMap(tasks []models.Task) map[int]int {
	idToIndex := make(map[int]int)

	for index, task := range tasks {
		fmt.Println(task.Id, task.Description)
		idToIndex[task.Id] = index
	}
	return  idToIndex
}
